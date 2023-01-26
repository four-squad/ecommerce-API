// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	cart "ecommerce/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartService is an autogenerated mock type for the CartService type
type CartService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, idProduct
func (_m *CartService) Add(token interface{}, idProduct uint) error {
	ret := _m.Called(token, idProduct)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, idProduct)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: token, idCart
func (_m *CartService) Delete(token interface{}, idCart uint) error {
	ret := _m.Called(token, idCart)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, idCart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByIdC provides a mock function with given fields: token, idCart
func (_m *CartService) GetByIdC(token interface{}, idCart uint) (cart.CoreCart, error) {
	ret := _m.Called(token, idCart)

	var r0 cart.CoreCart
	if rf, ok := ret.Get(0).(func(interface{}, uint) cart.CoreCart); ok {
		r0 = rf(token, idCart)
	} else {
		r0 = ret.Get(0).(cart.CoreCart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint) error); ok {
		r1 = rf(token, idCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdU provides a mock function with given fields: token
func (_m *CartService) GetByIdU(token interface{}) ([]cart.CoreCart, error) {
	ret := _m.Called(token)

	var r0 []cart.CoreCart
	if rf, ok := ret.Get(0).(func(interface{}) []cart.CoreCart); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.CoreCart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, idCart, updatedData
func (_m *CartService) Update(token interface{}, idCart uint, updatedData cart.CoreCart) (cart.CoreCart, error) {
	ret := _m.Called(token, idCart, updatedData)

	var r0 cart.CoreCart
	if rf, ok := ret.Get(0).(func(interface{}, uint, cart.CoreCart) cart.CoreCart); ok {
		r0 = rf(token, idCart, updatedData)
	} else {
		r0 = ret.Get(0).(cart.CoreCart)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, cart.CoreCart) error); ok {
		r1 = rf(token, idCart, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartService creates a new instance of CartService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartService(t mockConstructorTestingTNewCartService) *CartService {
	mock := &CartService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
