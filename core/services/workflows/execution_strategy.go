package workflows

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"

	"github.com/smartcontractkit/libocr/permutation"

	"golang.org/x/crypto/sha3"
)

type executionStrategy interface {
	Apply(ctx context.Context, l logger.Logger, cap capabilities.CallbackCapability, req capabilities.CapabilityRequest) (values.Value, error)
}

var _ executionStrategy = immediateExecution{}

type immediateExecution struct{}

func (i immediateExecution) Apply(ctx context.Context, lggr logger.Logger, cap capabilities.CallbackCapability, req capabilities.CapabilityRequest) (values.Value, error) {
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

var _ executionStrategy = scheduledExecution{}

type scheduledExecution struct {
	DON      *capabilities.DON
	PeerID   p2ptypes.PeerID
	Position int
}

var (
	// S = [N]
	Schedule_AllAtOnce = "allAtOnce"
	// S = [1 * N]
	Schedule_OneAtATime = "oneAtATime"
)

func (d scheduledExecution) Apply(ctx context.Context, lggr logger.Logger, cap capabilities.CallbackCapability, req capabilities.CapabilityRequest) (values.Value, error) {
	tc, err := d.transmissionConfig(req.Config)
	if err != nil {
		return nil, err
	}

	info, err := cap.Info(ctx)
	if err != nil {
		return nil, err
	}

	n := len(d.DON.Members)
	key := d.key(d.DON.Config.SharedSecret, req.Metadata.WorkflowID, req.Metadata.WorkflowExecutionID)
	position := d.Position

	// Note: if donInfo == nil, then this means the capability
	// is local. We'll use the local DON info passed into the engine
	// to generate the schedule.
	if info.DON != nil {
		n = len(info.DON.Members)
		key = d.key(info.DON.Config.SharedSecret, req.Metadata.WorkflowID, req.Metadata.WorkflowExecutionID)

		// Get our position in the target DON
		shuffledTargetDONPositions := permutation.Permutation(n, key)
		if position > n-1 {
			lggr.Debugw("skipping transmission: node is not included in schedule")
			return nil, nil
		}

		position = shuffledTargetDONPositions[position]
	}

	picked := permutation.Permutation(n, key)
	delay := d.delayFor(position, tc.Schedule, picked, tc.DeltaStage)
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

func (d scheduledExecution) key(sharedSecret [16]byte, workflowID, workflowExecutionID string) [16]byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(sharedSecret[:])
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

func (d scheduledExecution) transmissionConfig(config *values.Map) (transmissionConfig, error) {
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

	sched, err := schedule(tc.Schedule, len(d.DON.Members))
	if err != nil {
		return transmissionConfig{}, err
	}

	return transmissionConfig{
		Schedule:   sched,
		DeltaStage: duration,
	}, nil
}

func (d scheduledExecution) delayFor(position int, schedule []int, permutation []int, deltaStage time.Duration) *time.Duration {
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
