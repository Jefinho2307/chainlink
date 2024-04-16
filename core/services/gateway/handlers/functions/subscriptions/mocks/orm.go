// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	subscriptions "github.com/smartcontractkit/chainlink/v2/core/services/gateway/handlers/functions/subscriptions"
	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"
	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// GetSubscriptions provides a mock function with given fields: offset, limit, qopts
func (_m *ORM) GetSubscriptions(offset uint, limit uint, qopts ...pg.QOpt) ([]subscriptions.StoredSubscription, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSubscriptions")
	}

	var r0 []subscriptions.StoredSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint, ...pg.QOpt) ([]subscriptions.StoredSubscription, error)); ok {
		return rf(offset, limit, qopts...)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, ...pg.QOpt) []subscriptions.StoredSubscription); ok {
		r0 = rf(offset, limit, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subscriptions.StoredSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, uint, ...pg.QOpt) error); ok {
		r1 = rf(offset, limit, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertSubscription provides a mock function with given fields: subscription, qopts
func (_m *ORM) UpsertSubscription(subscription subscriptions.StoredSubscription, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, subscription)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpsertSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(subscriptions.StoredSubscription, ...pg.QOpt) error); ok {
		r0 = rf(subscription, qopts...)
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
