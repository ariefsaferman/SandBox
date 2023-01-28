// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// GetListHistoryTransaction provides a mock function with given fields: sender, search, sortBy, sort, limit
func (_m *TransactionRepository) GetListHistoryTransaction(sender uint, search string, sortBy string, sort string, limit string) ([]*entity.Transaction, error) {
	ret := _m.Called(sender, search, sortBy, sort, limit)

	var r0 []*entity.Transaction
	if rf, ok := ret.Get(0).(func(uint, string, string, string, string) []*entity.Transaction); ok {
		r0 = rf(sender, search, sortBy, sort, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, string, string, string, string) error); ok {
		r1 = rf(sender, search, sortBy, sort, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TopUp provides a mock function with given fields: tr, IdLogin
func (_m *TransactionRepository) TopUp(tr *entity.Transaction, IdLogin uint) (*entity.Transaction, error) {
	ret := _m.Called(tr, IdLogin)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(*entity.Transaction, uint) *entity.Transaction); ok {
		r0 = rf(tr, IdLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Transaction, uint) error); ok {
		r1 = rf(tr, IdLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: tr, sender, receiver
func (_m *TransactionRepository) Transfer(tr *entity.Transaction, sender uint, receiver uint) (*entity.Transaction, error) {
	ret := _m.Called(tr, sender, receiver)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(*entity.Transaction, uint, uint) *entity.Transaction); ok {
		r0 = rf(tr, sender, receiver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Transaction, uint, uint) error); ok {
		r1 = rf(tr, sender, receiver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
