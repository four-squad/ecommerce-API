// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Transaction_DetailHandler is an autogenerated mock type for the Transaction_DetailHandler type
type Transaction_DetailHandler struct {
	mock.Mock
}

type mockConstructorTestingTNewTransaction_DetailHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransaction_DetailHandler creates a new instance of Transaction_DetailHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransaction_DetailHandler(t mockConstructorTestingTNewTransaction_DetailHandler) *Transaction_DetailHandler {
	mock := &Transaction_DetailHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
