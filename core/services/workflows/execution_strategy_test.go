package workflows

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-common/pkg/values"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func assertBetween(t *testing.T, got time.Duration, low time.Duration, high time.Duration) {
	assert.GreaterOrEqual(t, got, low)
	assert.LessOrEqual(t, got, high)
}

func TestDelayedExecutionStrategy(t *testing.T) {
	var gotTime time.Time
	var called bool
	mt := newMockCapability(
		capabilities.MustNewCapabilityInfo(
			"write_polygon-testnet-mumbai",
			capabilities.CapabilityTypeTarget,
			"a write capability targeting polygon mumbai testnet",
			"v1.0.0",
		),
		func(req capabilities.CapabilityRequest) (capabilities.CapabilityResponse, error) {
			gotTime = time.Now()
			called = true
			return capabilities.CapabilityResponse{}, nil
		},
	)

	l := logger.TestLogger(t)

	// The combination of this key and the metadata above
	// will yield the permutation [3, 2, 0, 1]
	key, err := hex.DecodeString("fb13ca015a9ec60089c7141e9522de79")
	require.NoError(t, err)

	testCases := []struct {
		name     string
		position int
		schedule string
		low      time.Duration
		high     time.Duration
	}{
		{
			name:     "position 0; oneAtATime",
			position: 0,
			schedule: "oneAtATime",
			low:      300 * time.Millisecond,
			high:     400 * time.Millisecond,
		},
		{
			name:     "position 1; oneAtATime",
			position: 1,
			schedule: "oneAtATime",
			low:      200 * time.Millisecond,
			high:     300 * time.Millisecond,
		},
		{
			name:     "position 2; oneAtATime",
			position: 2,
			schedule: "oneAtATime",
			low:      0 * time.Millisecond,
			high:     100 * time.Millisecond,
		},
		{
			name:     "position 3; oneAtATime",
			position: 3,
			schedule: "oneAtATime",
			low:      100 * time.Millisecond,
			high:     200 * time.Millisecond,
		},
		{
			name:     "position 0; allAtOnce",
			position: 0,
			schedule: "allAtOnce",
			low:      0 * time.Millisecond,
			high:     100 * time.Millisecond,
		},
		{
			name:     "position 1; allAtOnce",
			position: 1,
			schedule: "allAtOnce",
			low:      0 * time.Millisecond,
			high:     100 * time.Millisecond,
		},
		{
			name:     "position 2; allAtOnce",
			position: 2,
			schedule: "allAtOnce",
			low:      0 * time.Millisecond,
			high:     100 * time.Millisecond,
		},
		{
			name:     "position 3; allAtOnce",
			position: 3,
			schedule: "allAtOnce",
			low:      0 * time.Millisecond,
			high:     100 * time.Millisecond,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			startTime := time.Now()

			m, err := values.NewMap(map[string]any{
				"schedule":   tc.schedule,
				"deltaStage": "100ms",
			})
			require.NoError(t, err)

			req := capabilities.CapabilityRequest{
				Config: m,
				Metadata: capabilities.RequestMetadata{
					WorkflowID:          "mock-workflow-id",
					WorkflowExecutionID: "mock-execution-id",
				},
			}

			de := scheduledExecution{
				sharedSecret: [16]byte(key),
				position:     tc.position,
				n:            4,
			}
			_, err = de.Apply(tests.Context(t), l, mt, req)
			require.NoError(t, err)
			require.True(t, called)

			assertBetween(t, gotTime.Sub(startTime), tc.low, tc.high)
		})
	}
}
