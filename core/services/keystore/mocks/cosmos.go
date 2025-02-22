// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cosmoskey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/cosmoskey"

	mock "github.com/stretchr/testify/mock"
)

// Cosmos is an autogenerated mock type for the Cosmos type
type Cosmos struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, key
func (_m *Cosmos) Add(ctx context.Context, key cosmoskey.Key) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, cosmoskey.Key) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx
func (_m *Cosmos) Create(ctx context.Context) (cosmoskey.Key, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 cosmoskey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (cosmoskey.Key, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) cosmoskey.Key); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(cosmoskey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Cosmos) Delete(ctx context.Context, id string) (cosmoskey.Key, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 cosmoskey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (cosmoskey.Key, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) cosmoskey.Key); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(cosmoskey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields: ctx
func (_m *Cosmos) EnsureKey(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *Cosmos) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Cosmos) Get(id string) (cosmoskey.Key, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 cosmoskey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (cosmoskey.Key, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) cosmoskey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(cosmoskey.Key)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Cosmos) GetAll() ([]cosmoskey.Key, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []cosmoskey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]cosmoskey.Key, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []cosmoskey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cosmoskey.Key)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *Cosmos) Import(ctx context.Context, keyJSON []byte, password string) (cosmoskey.Key, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 cosmoskey.Key
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (cosmoskey.Key, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) cosmoskey.Key); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(cosmoskey.Key)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCosmos creates a new instance of Cosmos. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCosmos(t interface {
	mock.TestingT
	Cleanup(func())
}) *Cosmos {
	mock := &Cosmos{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
