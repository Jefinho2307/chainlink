// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	big "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"

	mock "github.com/stretchr/testify/mock"

	s4 "github.com/smartcontractkit/chainlink/v2/core/services/s4"

	time "time"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// DeleteExpired provides a mock function with given fields: ctx, limit, utcNow
func (_m *ORM) DeleteExpired(ctx context.Context, limit uint, utcNow time.Time) (int64, error) {
	ret := _m.Called(ctx, limit, utcNow)

	if len(ret) == 0 {
		panic("no return value specified for DeleteExpired")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, time.Time) (int64, error)); ok {
		return rf(ctx, limit, utcNow)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint, time.Time) int64); ok {
		r0 = rf(ctx, limit, utcNow)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint, time.Time) error); ok {
		r1 = rf(ctx, limit, utcNow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, address, slotId
func (_m *ORM) Get(ctx context.Context, address *big.Big, slotId uint) (*s4.Row, error) {
	ret := _m.Called(ctx, address, slotId)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *s4.Row
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Big, uint) (*s4.Row, error)); ok {
		return rf(ctx, address, slotId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Big, uint) *s4.Row); ok {
		r0 = rf(ctx, address, slotId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s4.Row)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Big, uint) error); ok {
		r1 = rf(ctx, address, slotId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSnapshot provides a mock function with given fields: ctx, addressRange
func (_m *ORM) GetSnapshot(ctx context.Context, addressRange *s4.AddressRange) ([]*s4.SnapshotRow, error) {
	ret := _m.Called(ctx, addressRange)

	if len(ret) == 0 {
		panic("no return value specified for GetSnapshot")
	}

	var r0 []*s4.SnapshotRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s4.AddressRange) ([]*s4.SnapshotRow, error)); ok {
		return rf(ctx, addressRange)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s4.AddressRange) []*s4.SnapshotRow); ok {
		r0 = rf(ctx, addressRange)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*s4.SnapshotRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s4.AddressRange) error); ok {
		r1 = rf(ctx, addressRange)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnconfirmedRows provides a mock function with given fields: ctx, limit
func (_m *ORM) GetUnconfirmedRows(ctx context.Context, limit uint) ([]*s4.Row, error) {
	ret := _m.Called(ctx, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetUnconfirmedRows")
	}

	var r0 []*s4.Row
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]*s4.Row, error)); ok {
		return rf(ctx, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*s4.Row); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*s4.Row)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, row
func (_m *ORM) Update(ctx context.Context, row *s4.Row) error {
	ret := _m.Called(ctx, row)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *s4.Row) error); ok {
		r0 = rf(ctx, row)
	} else {
		r0 = ret.Error(0)
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
