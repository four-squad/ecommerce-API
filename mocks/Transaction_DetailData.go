// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Transaction_DetailData is an autogenerated mock type for the Transaction_DetailData type
type Transaction_DetailData struct {
	mock.Mock
}

type mockConstructorTestingTNewTransaction_DetailData interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransaction_DetailData creates a new instance of Transaction_DetailData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction_DetailData(t mockConstructorTestingTNewTransaction_DetailData) *Transaction_DetailData {
	mock := &Transaction_DetailData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
