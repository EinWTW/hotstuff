// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/relab/hotstuff (interfaces: Replica)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	hotstuff "github.com/relab/hotstuff"
	reflect "reflect"
)

// MockReplica is a mock of Replica interface
type MockReplica struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaMockRecorder
}

// MockReplicaMockRecorder is the mock recorder for MockReplica
type MockReplicaMockRecorder struct {
	mock *MockReplica
}

// NewMockReplica creates a new mock instance
func NewMockReplica(ctrl *gomock.Controller) *MockReplica {
	mock := &MockReplica{ctrl: ctrl}
	mock.recorder = &MockReplicaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReplica) EXPECT() *MockReplicaMockRecorder {
	return m.recorder
}

// Deliver mocks base method
func (m *MockReplica) Deliver(arg0 *hotstuff.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Deliver", arg0)
}

// Deliver indicates an expected call of Deliver
func (mr *MockReplicaMockRecorder) Deliver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deliver", reflect.TypeOf((*MockReplica)(nil).Deliver), arg0)
}

// ID mocks base method
func (m *MockReplica) ID() hotstuff.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(hotstuff.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockReplicaMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockReplica)(nil).ID))
}

// NewView mocks base method
func (m *MockReplica) NewView(arg0 hotstuff.QuorumCert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewView", arg0)
}

// NewView indicates an expected call of NewView
func (mr *MockReplicaMockRecorder) NewView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewView", reflect.TypeOf((*MockReplica)(nil).NewView), arg0)
}

// PublicKey mocks base method
func (m *MockReplica) PublicKey() hotstuff.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicKey")
	ret0, _ := ret[0].(hotstuff.PublicKey)
	return ret0
}

// PublicKey indicates an expected call of PublicKey
func (mr *MockReplicaMockRecorder) PublicKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicKey", reflect.TypeOf((*MockReplica)(nil).PublicKey))
}

// Vote mocks base method
func (m *MockReplica) Vote(arg0 hotstuff.PartialCert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Vote", arg0)
}

// Vote indicates an expected call of Vote
func (mr *MockReplicaMockRecorder) Vote(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockReplica)(nil).Vote), arg0)
}
