// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Eth is an autogenerated mock type for the Eth type
type Eth struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, address, chainID, qopts
func (_m *Eth) Add(ctx context.Context, address common.Address, chainID *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(ctx, address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEnabled provides a mock function with given fields: ctx, address, chainID
func (_m *Eth) CheckEnabled(ctx context.Context, address common.Address, chainID *big.Int) error {
	ret := _m.Called(ctx, address, chainID)

	if len(ret) == 0 {
		panic("no return value specified for CheckEnabled")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) error); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, chainIDs
func (_m *Eth) Create(ctx context.Context, chainIDs ...*big.Int) (ethkey.KeyV2, error) {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*big.Int) (ethkey.KeyV2, error)); ok {
		return rf(ctx, chainIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*big.Int) ethkey.KeyV2); ok {
		r0 = rf(ctx, chainIDs...)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*big.Int) error); ok {
		r1 = rf(ctx, chainIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Eth) Delete(ctx context.Context, id string) (ethkey.KeyV2, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (ethkey.KeyV2, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) ethkey.KeyV2); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disable provides a mock function with given fields: ctx, address, chainID, qopts
func (_m *Eth) Disable(ctx context.Context, address common.Address, chainID *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Disable")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(ctx, address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Enable provides a mock function with given fields: ctx, address, chainID, qopts
func (_m *Eth) Enable(ctx context.Context, address common.Address, chainID *big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Enable")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, ...pg.QOpt) error); ok {
		r0 = rf(ctx, address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnabledAddressesForChain provides a mock function with given fields: ctx, chainID
func (_m *Eth) EnabledAddressesForChain(ctx context.Context, chainID *big.Int) ([]common.Address, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for EnabledAddressesForChain")
	}

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) ([]common.Address, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) []common.Address); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnabledKeysForChain provides a mock function with given fields: ctx, chainID
func (_m *Eth) EnabledKeysForChain(ctx context.Context, chainID *big.Int) ([]ethkey.KeyV2, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for EnabledKeysForChain")
	}

	var r0 []ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) ([]ethkey.KeyV2, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) []ethkey.KeyV2); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKeys provides a mock function with given fields: ctx, chainIDs
func (_m *Eth) EnsureKeys(ctx context.Context, chainIDs ...*big.Int) error {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EnsureKeys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*big.Int) error); ok {
		r0 = rf(ctx, chainIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: ctx, id, password
func (_m *Eth) Export(ctx context.Context, id string, password string) ([]byte, error) {
	ret := _m.Called(ctx, id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]byte, error)); ok {
		return rf(ctx, id, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []byte); ok {
		r0 = rf(ctx, id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *Eth) Get(ctx context.Context, id string) (ethkey.KeyV2, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (ethkey.KeyV2, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) ethkey.KeyV2); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Eth) GetAll(ctx context.Context) ([]ethkey.KeyV2, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]ethkey.KeyV2, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []ethkey.KeyV2); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoundRobinAddress provides a mock function with given fields: ctx, chainID, addresses
func (_m *Eth) GetRoundRobinAddress(ctx context.Context, chainID *big.Int, addresses ...common.Address) (common.Address, error) {
	_va := make([]interface{}, len(addresses))
	for _i := range addresses {
		_va[_i] = addresses[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRoundRobinAddress")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, ...common.Address) (common.Address, error)); ok {
		return rf(ctx, chainID, addresses...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, ...common.Address) common.Address); ok {
		r0 = rf(ctx, chainID, addresses...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, ...common.Address) error); ok {
		r1 = rf(ctx, chainID, addresses...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: ctx, id, chainID
func (_m *Eth) GetState(ctx context.Context, id string, chainID *big.Int) (ethkey.State, error) {
	ret := _m.Called(ctx, id, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *big.Int) (ethkey.State, error)); ok {
		return rf(ctx, id, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *big.Int) ethkey.State); ok {
		r0 = rf(ctx, id, chainID)
	} else {
		r0 = ret.Get(0).(ethkey.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *big.Int) error); ok {
		r1 = rf(ctx, id, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateForKey provides a mock function with given fields: ctx, key
func (_m *Eth) GetStateForKey(ctx context.Context, key ethkey.KeyV2) (ethkey.State, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetStateForKey")
	}

	var r0 ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethkey.KeyV2) (ethkey.State, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethkey.KeyV2) ethkey.State); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(ethkey.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethkey.KeyV2) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatesForChain provides a mock function with given fields: ctx, chainID
func (_m *Eth) GetStatesForChain(ctx context.Context, chainID *big.Int) ([]ethkey.State, error) {
	ret := _m.Called(ctx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GetStatesForChain")
	}

	var r0 []ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) ([]ethkey.State, error)); ok {
		return rf(ctx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) []ethkey.State); ok {
		r0 = rf(ctx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.State)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatesForKeys provides a mock function with given fields: ctx, keys
func (_m *Eth) GetStatesForKeys(ctx context.Context, keys []ethkey.KeyV2) ([]ethkey.State, error) {
	ret := _m.Called(ctx, keys)

	if len(ret) == 0 {
		panic("no return value specified for GetStatesForKeys")
	}

	var r0 []ethkey.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []ethkey.KeyV2) ([]ethkey.State, error)); ok {
		return rf(ctx, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []ethkey.KeyV2) []ethkey.State); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ethkey.State)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []ethkey.KeyV2) error); ok {
		r1 = rf(ctx, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: ctx, keyJSON, password, chainIDs
func (_m *Eth) Import(ctx context.Context, keyJSON []byte, password string, chainIDs ...*big.Int) (ethkey.KeyV2, error) {
	_va := make([]interface{}, len(chainIDs))
	for _i := range chainIDs {
		_va[_i] = chainIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, keyJSON, password)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 ethkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string, ...*big.Int) (ethkey.KeyV2, error)); ok {
		return rf(ctx, keyJSON, password, chainIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string, ...*big.Int) ethkey.KeyV2); ok {
		r0 = rf(ctx, keyJSON, password, chainIDs...)
	} else {
		r0 = ret.Get(0).(ethkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string, ...*big.Int) error); ok {
		r1 = rf(ctx, keyJSON, password, chainIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTx provides a mock function with given fields: ctx, fromAddress, tx, chainID
func (_m *Eth) SignTx(ctx context.Context, fromAddress common.Address, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(ctx, fromAddress, tx, chainID)

	if len(ret) == 0 {
		panic("no return value specified for SignTx")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction, *big.Int) (*types.Transaction, error)); ok {
		return rf(ctx, fromAddress, tx, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(ctx, fromAddress, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *types.Transaction, *big.Int) error); ok {
		r1 = rf(ctx, fromAddress, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToKeyChanges provides a mock function with given fields: ctx
func (_m *Eth) SubscribeToKeyChanges(ctx context.Context) (chan struct{}, func()) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeToKeyChanges")
	}

	var r0 chan struct{}
	var r1 func()
	if rf, ok := ret.Get(0).(func(context.Context) (chan struct{}, func())); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) chan struct{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) func()); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}

// XXXTestingOnlyAdd provides a mock function with given fields: ctx, key
func (_m *Eth) XXXTestingOnlyAdd(ctx context.Context, key ethkey.KeyV2) {
	_m.Called(ctx, key)
}

// XXXTestingOnlySetState provides a mock function with given fields: ctx, keyState
func (_m *Eth) XXXTestingOnlySetState(ctx context.Context, keyState ethkey.State) {
	_m.Called(ctx, keyState)
}

// NewEth creates a new instance of Eth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEth(t interface {
	mock.TestingT
	Cleanup(func())
}) *Eth {
	mock := &Eth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
