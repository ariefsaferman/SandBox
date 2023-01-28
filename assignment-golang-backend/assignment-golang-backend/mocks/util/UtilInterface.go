// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UtilInterface is an autogenerated mock type for the UtilInterface type
type UtilInterface struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashedPwd, inputPwd
func (_m *UtilInterface) ComparePassword(hashedPwd string, inputPwd string) bool {
	ret := _m.Called(hashedPwd, inputPwd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPwd, inputPwd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewUtilInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUtilInterface creates a new instance of UtilInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUtilInterface(t mockConstructorTestingTNewUtilInterface) *UtilInterface {
	mock := &UtilInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
