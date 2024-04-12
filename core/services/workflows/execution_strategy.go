package workflows

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink/v2/core/logger"

	"github.com/smartcontractkit/libocr/permutation"

	"golang.org/x/crypto/sha3"
)

type executionStrategy interface {
	Apply(ctx context.Context, l logger.Logger, cap capabilities.CallbackExecutable, req capabilities.CapabilityRequest) (values.Value, error)
}

var _ executionStrategy = immediateExecution{}

type immediateExecution struct{}

func (i immediateExecution) Apply(ctx context.Context, lggr logger.Logger, cap capabilities.CallbackExecutable, req capabilities.CapabilityRequest) (values.Value, error) {
	l, err := capabilities.ExecuteSync(ctx, cap, req)
	if err != nil {
		return nil, err
	}

	// `ExecuteSync` returns a `values.List` even if there was
	// just one return value. If that is the case, let's unwrap the
	// single value to make it easier to use in -- for example -- variable interpolation.
	if len(l.Underlying) > 1 {
		return l, nil
	}

	return l.Underlying[0], nil
}

var _ executionStrategy = delayedExecution{}

type delayedExecution struct {
	sharedSecret [16]byte
	// Position of the current node in the list of DON nodes maintained by the Registry
	position int
	// Number of nodes in the workflow DON
	n int
}

var (
	// S = [N]
	Schedule_AllAtOnce = "allAtOnce"
	// S = [1 * N]
	Schedule_OneAtATime = "oneAtATime"
)

func (d delayedExecution) Apply(ctx context.Context, lggr logger.Logger, cap capabilities.CallbackExecutable, req capabilities.CapabilityRequest) (values.Value, error) {
	tc, err := d.transmissionConfig(req.Config)
	if err != nil {
		return nil, err
	}

	md := req.Metadata
	key := d.key(md.WorkflowID, md.WorkflowExecutionID)
	picked := permutation.Permutation(d.n, key)

	delay := d.delayFor(d.position, tc.Schedule, picked, tc.DeltaStage)
	if delay == nil {
		lggr.Debugw("skipping transmission: node is not included in schedule")
		return nil, nil
	}

	lggr.Debugf("execution delayed by %+v", *delay)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(*delay):
		lggr.Debugw("executing delayed execution")
		return immediateExecution{}.Apply(ctx, lggr, cap, req)
	}
}

func (d delayedExecution) key(workflowID, workflowExecutionID string) [16]byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(d.sharedSecret[:])
	hash.Write([]byte(workflowID))
	hash.Write([]byte(workflowExecutionID))

	var key [16]byte
	copy(key[:], hash.Sum(nil))
	return key
}

type transmissionConfig struct {
	Schedule   []int
	DeltaStage time.Duration
}

func (d delayedExecution) transmissionConfig(config *values.Map) (transmissionConfig, error) {
	var tc struct {
		DeltaStage string
		Schedule   string
	}
	err := config.UnwrapTo(&tc)
	if err != nil {
		return transmissionConfig{}, err
	}

	duration, err := time.ParseDuration(tc.DeltaStage)
	if err != nil {
		return transmissionConfig{}, fmt.Errorf("failed to parse DeltaStage %s as duration: %w", tc.DeltaStage, err)
	}

	sched, err := schedule(tc.Schedule, d.n)
	if err != nil {
		return transmissionConfig{}, err
	}

	return transmissionConfig{
		Schedule:   sched,
		DeltaStage: duration,
	}, nil
}

func (d delayedExecution) delayFor(position int, schedule []int, permutation []int, deltaStage time.Duration) *time.Duration {
	sum := 0
	for i, s := range schedule {
		sum += s
		if permutation[position] < sum {
			result := time.Duration(i) * deltaStage
			return &result
		}
	}

	return nil
}

func schedule(sched string, N int) ([]int, error) {
	switch sched {
	case Schedule_AllAtOnce:
		return []int{N}, nil
	case Schedule_OneAtATime:
		sch := []int{}
		for i := 0; i < N; i++ {
			sch = append(sch, 1)
		}
		return sch, nil
	}
	return nil, fmt.Errorf("unknown schedule %s", sched)
}
