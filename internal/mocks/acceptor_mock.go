// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/relab/hotstuff/modules (interfaces: Acceptor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hotstuff "github.com/relab/hotstuff"
)

// MockAcceptor is a mock of Acceptor interface.
type MockAcceptor struct {
	ctrl     *gomock.Controller
	recorder *MockAcceptorMockRecorder
}

// MockAcceptorMockRecorder is the mock recorder for MockAcceptor.
type MockAcceptorMockRecorder struct {
	mock *MockAcceptor
}

// NewMockAcceptor creates a new mock instance.
func NewMockAcceptor(ctrl *gomock.Controller) *MockAcceptor {
	mock := &MockAcceptor{ctrl: ctrl}
	mock.recorder = &MockAcceptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAcceptor) EXPECT() *MockAcceptorMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockAcceptor) Accept(arg0 hotstuff.Command) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockAcceptorMockRecorder) Accept(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockAcceptor)(nil).Accept), arg0)
}

// Proposed mocks base method.
func (m *MockAcceptor) Proposed(arg0 hotstuff.Command) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Proposed", arg0)
}

// Proposed indicates an expected call of Proposed.
func (mr *MockAcceptorMockRecorder) Proposed(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proposed", reflect.TypeOf((*MockAcceptor)(nil).Proposed), arg0)
}
