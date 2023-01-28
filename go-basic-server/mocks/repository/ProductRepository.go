// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	entity "go-basic-server/entity"

	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: p
func (_m *ProductRepository) CreateProduct(p *entity.Product) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Product) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProduct provides a mock function with given fields: id
func (_m *ProductRepository) DeleteProduct(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProduct provides a mock function with given fields:
func (_m *ProductRepository) GetAllProduct() []*entity.Product {
	ret := _m.Called()

	var r0 []*entity.Product
	if rf, ok := ret.Get(0).(func() []*entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Product)
		}
	}

	return r0
}

// GetOneProductById provides a mock function with given fields: id
func (_m *ProductRepository) GetOneProductById(id int) (*entity.Product, error) {
	ret := _m.Called(id)

	var r0 *entity.Product
	if rf, ok := ret.Get(0).(func(int) *entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
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

// UpdateProduct provides a mock function with given fields: id, p
func (_m *ProductRepository) UpdateProduct(id int, p *entity.Product) error {
	ret := _m.Called(id, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *entity.Product) error); ok {
		r0 = rf(id, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
