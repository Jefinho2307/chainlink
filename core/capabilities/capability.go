package capabilities

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/mod/semver"

	"github.com/smartcontractkit/chainlink-common/pkg/values"
)

// CapabilityType is an enum for the type of capability.
type CapabilityType int

// CapabilityType enum values.
const (
	CapabilityTypeTrigger CapabilityType = iota
	CapabilityTypeAction
	CapabilityTypeConsensus
	CapabilityTypeTarget
)

// String returns a string representation of CapabilityType
func (c CapabilityType) String() string {
	switch c {
	case CapabilityTypeTrigger:
		return "trigger"
	case CapabilityTypeAction:
		return "action"
	case CapabilityTypeConsensus:
		return "report"
	case CapabilityTypeTarget:
		return "target"
	}

	// Panic as this should be unreachable.
	panic("unknown capability type")
}

// IsValid checks if the capability type is valid.
func (c CapabilityType) IsValid() error {
	switch c {
	case CapabilityTypeTrigger,
		CapabilityTypeAction,
		CapabilityTypeConsensus,
		CapabilityTypeTarget:
		return nil
	}

	return fmt.Errorf("invalid capability type: %s", c)
}

// Validatable is an interface for validating the config and inputs of a capability.
type Validatable interface {
	ExampleOutput(inputs values.Map) values.Value
	ValidateInput(inputs values.Map) error
}

// CapabilityResponse is a struct for the Execute response of a capability.
type CapabilityResponse struct {
	value values.Value
	err   error
}

// Executable is an interface for executing a capability.
type Executable interface {
	// Start will be called when the capability is loaded by the application.
	// Start will be called before the capability is added to the registry.
	Start(ctx context.Context, config values.Map) (values.Value, error)
	// Capability must respect context.Done and cleanup any request specific resources
	// when the context is cancelled. When a request has been completed the capability
	// is also expected to close the callback channel.
	// Request specific configuration is passed in via the inputs parameter.
	// A successful response must always return a value. An error is assumed otherwise.
	// The intent is to make the API explicit.
	Execute(ctx context.Context, callback chan CapabilityResponse, inputs values.Map) error
	// Stop will be called before the application exits.
	// Stop will be called after the capability is removed from the registry.
	Stop(ctx context.Context) error
}

// Capability is an interface for a capability.
type Capability interface {
	Executable
	Validatable
	Info() CapabilityInfo
}

// CapabilityInfo is a struct for the info of a capability.
type CapabilityInfo struct {
	ID             string
	CapabilityType CapabilityType
	Description    string
	Version        string
}

// Info returns the info of the capability.
func (c CapabilityInfo) Info() CapabilityInfo {
	return c
}

var idRegex = regexp.MustCompile(`[a-z0-9_\-:]`)

const (
	idMaxLength = 128
)

// NewCapabilityInfo returns a new CapabilityInfo.
func NewCapabilityInfo(
	id string,
	capabilityType CapabilityType,
	description string,
	version string,
) (CapabilityInfo, error) {
	if len(id) > idMaxLength {
		return CapabilityInfo{}, fmt.Errorf("invalid id: %s exceeds max length %d", id, idMaxLength)
	}
	if !idRegex.MatchString(id) {
		return CapabilityInfo{}, fmt.Errorf("invalid id: %s. Allowed: %s", id, idRegex)
	}

	if ok := semver.IsValid(version); !ok {
		return CapabilityInfo{}, fmt.Errorf("invalid version: %+v", version)
	}

	if err := capabilityType.IsValid(); err != nil {
		return CapabilityInfo{}, err
	}

	return CapabilityInfo{
		ID:             id,
		CapabilityType: capabilityType,
		Description:    description,
		Version:        version,
	}, nil
}

var defaultExecuteTimeout = 10 * time.Second

// ExecuteSync executes a capability synchronously.
// We are not handling a case where a capability panics and crashes.
// There is default timeout of 10 seconds. If a capability takes longer than
// that then it should be executed asynchronously.
func ExecuteSync(ctx context.Context, c Capability, inputs values.Map) (values.Value, error) {
	ctxWithT, cancel := context.WithTimeout(ctx, defaultExecuteTimeout)
	defer cancel()

	callback := make(chan CapabilityResponse)
	vs := make([]values.Value, 0)

	var setupErr error
	go func(innerCtx context.Context, innerC Capability, innerInputs values.Map, innerCallback chan CapabilityResponse) {
		setupErr = innerC.Execute(innerCtx, innerCallback, innerInputs)
	}(ctxWithT, c, inputs, callback)

outerLoop:
	for {
		select {
		case response, isOpen := <-callback:
			if !isOpen {
				break outerLoop
			}
			// An error means execution has been interrupted.
			// We'll return the value discarding values received
			// until now.
			if response.err != nil {
				return nil, response.err
			}

			vs = append(vs, response.value)

		// Timeout when a capability panics, crashes, and does not close the channel.
		case <-ctxWithT.Done():
			return nil, fmt.Errorf("context timed out. If you did not set a timeout, be aware that the default ExecuteSync timeout is %f seconds", defaultExecuteTimeout.Seconds())
		}
	}

	// Something went wrong when setting up a capability.
	if setupErr != nil {
		return nil, setupErr
	}

	// If the capability did not return any values, we deem it as an error.
	// The intent is for the API to be explicit.
	if len(vs) == 0 {
		return nil, errors.New("capability did not return any values")
	}

	// If the capability returned only one value,
	// let's unwrap it to improve usability.
	if len(vs) == 1 {
		return vs[0], nil
	}

	return &values.List{Underlying: vs}, nil
}
