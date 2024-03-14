// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	big "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"

	context "context"

	job "github.com/smartcontractkit/chainlink/v2/core/services/job"

	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/smartcontractkit/chainlink/v2/core/services/pipeline"

	sqlutil "github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	types "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"

	uuid "github.com/google/uuid"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// AssertBridgesExist provides a mock function with given fields: ctx, p
func (_m *ORM) AssertBridgesExist(ctx context.Context, p pipeline.Pipeline) error {
	ret := _m.Called(ctx, p)

	if len(ret) == 0 {
		panic("no return value specified for AssertBridgesExist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pipeline.Pipeline) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountPipelineRunsByJobID provides a mock function with given fields: ctx, jobID
func (_m *ORM) CountPipelineRunsByJobID(ctx context.Context, jobID int32) (int32, error) {
	ret := _m.Called(ctx, jobID)

	if len(ret) == 0 {
		panic("no return value specified for CountPipelineRunsByJobID")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (int32, error)); ok {
		return rf(ctx, jobID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) int32); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJob provides a mock function with given fields: ctx, jb
func (_m *ORM) CreateJob(ctx context.Context, jb *job.Job) error {
	ret := _m.Called(ctx, jb)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.Job) error); ok {
		r0 = rf(ctx, jb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataSource provides a mock function with given fields:
func (_m *ORM) DataSource() sqlutil.DataSource {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DataSource")
	}

	var r0 sqlutil.DataSource
	if rf, ok := ret.Get(0).(func() sqlutil.DataSource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sqlutil.DataSource)
		}
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, id
func (_m *ORM) DeleteJob(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DismissError provides a mock function with given fields: ctx, errorID
func (_m *ORM) DismissError(ctx context.Context, errorID int64) error {
	ret := _m.Called(ctx, errorID)

	if len(ret) == 0 {
		panic("no return value specified for DismissError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, errorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindJob provides a mock function with given fields: ctx, id
func (_m *ORM) FindJob(ctx context.Context, id int32) (job.Job, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJob")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (job.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) job.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobByExternalJobID provides a mock function with given fields: ctx, _a1
func (_m *ORM) FindJobByExternalJobID(ctx context.Context, _a1 uuid.UUID) (job.Job, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for FindJobByExternalJobID")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (job.Job, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) job.Job); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobIDByAddress provides a mock function with given fields: ctx, address, evmChainID
func (_m *ORM) FindJobIDByAddress(ctx context.Context, address types.EIP55Address, evmChainID *big.Big) (int32, error) {
	ret := _m.Called(ctx, address, evmChainID)

	if len(ret) == 0 {
		panic("no return value specified for FindJobIDByAddress")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.EIP55Address, *big.Big) (int32, error)); ok {
		return rf(ctx, address, evmChainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.EIP55Address, *big.Big) int32); ok {
		r0 = rf(ctx, address, evmChainID)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.EIP55Address, *big.Big) error); ok {
		r1 = rf(ctx, address, evmChainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobIDsWithBridge provides a mock function with given fields: ctx, name
func (_m *ORM) FindJobIDsWithBridge(ctx context.Context, name string) ([]int32, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for FindJobIDsWithBridge")
	}

	var r0 []int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]int32, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []int32); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobTx provides a mock function with given fields: ctx, id
func (_m *ORM) FindJobTx(ctx context.Context, id int32) (job.Job, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJobTx")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (job.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) job.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobWithoutSpecErrors provides a mock function with given fields: ctx, id
func (_m *ORM) FindJobWithoutSpecErrors(ctx context.Context, id int32) (job.Job, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJobWithoutSpecErrors")
	}

	var r0 job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (job.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) job.Job); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.Job)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJobs provides a mock function with given fields: ctx, offset, limit
func (_m *ORM) FindJobs(ctx context.Context, offset int, limit int) ([]job.Job, int, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindJobs")
	}

	var r0 []job.Job
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]job.Job, int, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []job.Job); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindJobsByPipelineSpecIDs provides a mock function with given fields: ctx, ids
func (_m *ORM) FindJobsByPipelineSpecIDs(ctx context.Context, ids []int32) ([]job.Job, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for FindJobsByPipelineSpecIDs")
	}

	var r0 []job.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int32) ([]job.Job, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int32) []job.Job); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int32) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOCR2JobIDByAddress provides a mock function with given fields: ctx, contractID, feedID
func (_m *ORM) FindOCR2JobIDByAddress(ctx context.Context, contractID string, feedID *common.Hash) (int32, error) {
	ret := _m.Called(ctx, contractID, feedID)

	if len(ret) == 0 {
		panic("no return value specified for FindOCR2JobIDByAddress")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *common.Hash) (int32, error)); ok {
		return rf(ctx, contractID, feedID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *common.Hash) int32); ok {
		r0 = rf(ctx, contractID, feedID)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *common.Hash) error); ok {
		r1 = rf(ctx, contractID, feedID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunByID provides a mock function with given fields: ctx, id
func (_m *ORM) FindPipelineRunByID(ctx context.Context, id int64) (pipeline.Run, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunByID")
	}

	var r0 pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (pipeline.Run, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) pipeline.Run); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pipeline.Run)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunIDsByJobID provides a mock function with given fields: ctx, jobID, offset, limit
func (_m *ORM) FindPipelineRunIDsByJobID(ctx context.Context, jobID int32, offset int, limit int) ([]int64, error) {
	ret := _m.Called(ctx, jobID, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunIDsByJobID")
	}

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int, int) ([]int64, error)); ok {
		return rf(ctx, jobID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int, int) []int64); ok {
		r0 = rf(ctx, jobID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int, int) error); ok {
		r1 = rf(ctx, jobID, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPipelineRunsByIDs provides a mock function with given fields: ctx, ids
func (_m *ORM) FindPipelineRunsByIDs(ctx context.Context, ids []int64) ([]pipeline.Run, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for FindPipelineRunsByIDs")
	}

	var r0 []pipeline.Run
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64) ([]pipeline.Run, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int64) []pipeline.Run); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int64) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSpecError provides a mock function with given fields: ctx, id
func (_m *ORM) FindSpecError(ctx context.Context, id int64) (job.SpecError, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindSpecError")
	}

	var r0 job.SpecError
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (job.SpecError, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) job.SpecError); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(job.SpecError)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindSpecErrorsByJobIDs provides a mock function with given fields: ctx, ids
func (_m *ORM) FindSpecErrorsByJobIDs(ctx context.Context, ids []int32) ([]job.SpecError, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for FindSpecErrorsByJobIDs")
	}

	var r0 []job.SpecError
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []int32) ([]job.SpecError, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []int32) []job.SpecError); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.SpecError)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []int32) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTaskResultByRunIDAndTaskName provides a mock function with given fields: ctx, runID, taskName
func (_m *ORM) FindTaskResultByRunIDAndTaskName(ctx context.Context, runID int64, taskName string) ([]byte, error) {
	ret := _m.Called(ctx, runID, taskName)

	if len(ret) == 0 {
		panic("no return value specified for FindTaskResultByRunIDAndTaskName")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) ([]byte, error)); ok {
		return rf(ctx, runID, taskName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) []byte); ok {
		r0 = rf(ctx, runID, taskName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, runID, taskName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertJob provides a mock function with given fields: ctx, _a1
func (_m *ORM) InsertJob(ctx context.Context, _a1 *job.Job) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for InsertJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.Job) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertWebhookSpec provides a mock function with given fields: ctx, webhookSpec
func (_m *ORM) InsertWebhookSpec(ctx context.Context, webhookSpec *job.WebhookSpec) error {
	ret := _m.Called(ctx, webhookSpec)

	if len(ret) == 0 {
		panic("no return value specified for InsertWebhookSpec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.WebhookSpec) error); ok {
		r0 = rf(ctx, webhookSpec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PipelineRuns provides a mock function with given fields: ctx, jobID, offset, size
func (_m *ORM) PipelineRuns(ctx context.Context, jobID *int32, offset int, size int) ([]pipeline.Run, int, error) {
	ret := _m.Called(ctx, jobID, offset, size)

	if len(ret) == 0 {
		panic("no return value specified for PipelineRuns")
	}

	var r0 []pipeline.Run
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *int32, int, int) ([]pipeline.Run, int, error)); ok {
		return rf(ctx, jobID, offset, size)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int32, int, int) []pipeline.Run); ok {
		r0 = rf(ctx, jobID, offset, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int32, int, int) int); ok {
		r1 = rf(ctx, jobID, offset, size)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *int32, int, int) error); ok {
		r2 = rf(ctx, jobID, offset, size)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RecordError provides a mock function with given fields: ctx, jobID, description
func (_m *ORM) RecordError(ctx context.Context, jobID int32, description string) error {
	ret := _m.Called(ctx, jobID, description)

	if len(ret) == 0 {
		panic("no return value specified for RecordError")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, string) error); ok {
		r0 = rf(ctx, jobID, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TryRecordError provides a mock function with given fields: ctx, jobID, description
func (_m *ORM) TryRecordError(ctx context.Context, jobID int32, description string) {
	_m.Called(ctx, jobID, description)
}

// WithDataSource provides a mock function with given fields: source
func (_m *ORM) WithDataSource(source sqlutil.DataSource) job.ORM {
	ret := _m.Called(source)

	if len(ret) == 0 {
		panic("no return value specified for WithDataSource")
	}

	var r0 job.ORM
	if rf, ok := ret.Get(0).(func(sqlutil.DataSource) job.ORM); ok {
		r0 = rf(source)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(job.ORM)
		}
	}

	return r0
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
