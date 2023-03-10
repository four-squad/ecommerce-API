// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	product "ecommerce/features/product"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// Add provides a mock function with given fields: newContent, token, image
func (_m *ProductService) Add(newContent product.CoreProduct, token interface{}, image *multipart.FileHeader) (product.CoreProduct, error) {
	ret := _m.Called(newContent, token, image)

	var r0 product.CoreProduct
	if rf, ok := ret.Get(0).(func(product.CoreProduct, interface{}, *multipart.FileHeader) product.CoreProduct); ok {
		r0 = rf(newContent, token, image)
	} else {
		r0 = ret.Get(0).(product.CoreProduct)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(product.CoreProduct, interface{}, *multipart.FileHeader) error); ok {
		r1 = rf(newContent, token, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, contentId
func (_m *ProductService) Delete(token interface{}, contentId uint) error {
	ret := _m.Called(token, contentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, contentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ProductService) GetAll() ([]product.CoreProduct, error) {
	ret := _m.Called()

	var r0 []product.CoreProduct
	if rf, ok := ret.Get(0).(func() []product.CoreProduct); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.CoreProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: tes
func (_m *ProductService) GetById(tes uint) ([]product.CoreProduct, error) {
	ret := _m.Called(tes)

	var r0 []product.CoreProduct
	if rf, ok := ret.Get(0).(func(uint) []product.CoreProduct); ok {
		r0 = rf(tes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.CoreProduct)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(tes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, id, updatedData, file
func (_m *ProductService) Update(token interface{}, id uint, updatedData product.CoreProduct, file *multipart.FileHeader) (product.CoreProduct, error) {
	ret := _m.Called(token, id, updatedData, file)

	var r0 product.CoreProduct
	if rf, ok := ret.Get(0).(func(interface{}, uint, product.CoreProduct, *multipart.FileHeader) product.CoreProduct); ok {
		r0 = rf(token, id, updatedData, file)
	} else {
		r0 = ret.Get(0).(product.CoreProduct)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, product.CoreProduct, *multipart.FileHeader) error); ok {
		r1 = rf(token, id, updatedData, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductService interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductService creates a new instance of ProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductService(t mockConstructorTestingTNewProductService) *ProductService {
	mock := &ProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
