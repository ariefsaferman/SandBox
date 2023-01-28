// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	dto "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/dto"
	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: _a0
func (_m *AuthUsecase) Login(_a0 dto.AuthRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(_a0)

	var r0 *dto.LoginResponse
	if rf, ok := ret.Get(0).(func(dto.AuthRequest) *dto.LoginResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.AuthRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *AuthUsecase) Register(_a0 dto.AuthRequest) (*dto.RegisterResponse, error) {
	ret := _m.Called(_a0)

	var r0 *dto.RegisterResponse
	if rf, ok := ret.Get(0).(func(dto.AuthRequest) *dto.RegisterResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.AuthRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUsecase(t mockConstructorTestingTNewAuthUsecase) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}