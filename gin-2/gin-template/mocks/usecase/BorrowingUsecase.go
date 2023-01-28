// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	mock "github.com/stretchr/testify/mock"
)

// BorrowingUsecase is an autogenerated mock type for the BorrowingUsecase type
type BorrowingUsecase struct {
	mock.Mock
}

// AddBorrowBook provides a mock function with given fields: book
func (_m *BorrowingUsecase) AddBorrowBook(book *entity.BorrowingRecord) (*entity.BorrowingRecord, error) {
	ret := _m.Called(book)

	var r0 *entity.BorrowingRecord
	if rf, ok := ret.Get(0).(func(*entity.BorrowingRecord) *entity.BorrowingRecord); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.BorrowingRecord) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRecordById provides a mock function with given fields: id
func (_m *BorrowingUsecase) GetRecordById(id int) (*entity.BorrowingRecord, error) {
	ret := _m.Called(id)

	var r0 *entity.BorrowingRecord
	if rf, ok := ret.Get(0).(func(int) *entity.BorrowingRecord); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnBorrowBook provides a mock function with given fields: id
func (_m *BorrowingUsecase) ReturnBorrowBook(id int) (*entity.BorrowingRecord, error) {
	ret := _m.Called(id)

	var r0 *entity.BorrowingRecord
	if rf, ok := ret.Get(0).(func(int) *entity.BorrowingRecord); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecord)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowingUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowingUsecase creates a new instance of BorrowingUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowingUsecase(t mockConstructorTestingTNewBorrowingUsecase) *BorrowingUsecase {
	mock := &BorrowingUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}