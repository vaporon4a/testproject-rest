// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// BalanceChanger is an autogenerated mock type for the BalanceChanger type
type BalanceChanger struct {
	mock.Mock
}

// Deposit provides a mock function with given fields: ctx, walletUuid, amount
func (_m *BalanceChanger) Deposit(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	ret := _m.Called(ctx, walletUuid, amount)

	if len(ret) == 0 {
		panic("no return value specified for Deposit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int64) error); ok {
		r0 = rf(ctx, walletUuid, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Withdraw provides a mock function with given fields: ctx, walletUuid, amount
func (_m *BalanceChanger) Withdraw(ctx context.Context, walletUuid uuid.UUID, amount int64) error {
	ret := _m.Called(ctx, walletUuid, amount)

	if len(ret) == 0 {
		panic("no return value specified for Withdraw")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int64) error); ok {
		r0 = rf(ctx, walletUuid, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBalanceChanger creates a new instance of BalanceChanger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBalanceChanger(t interface {
	mock.TestingT
	Cleanup(func())
}) *BalanceChanger {
	mock := &BalanceChanger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
