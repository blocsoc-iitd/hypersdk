// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AnomalyFi/hypersdk/chain (interfaces: Auth)

// Package chain is a generated GoMock package.
package chain

import (
	context "context"
	reflect "reflect"

	codec "github.com/AnomalyFi/hypersdk/codec"
	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// AsyncVerify mocks base method.
func (m *MockAuth) AsyncVerify(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncVerify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AsyncVerify indicates an expected call of AsyncVerify.
func (mr *MockAuthMockRecorder) AsyncVerify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncVerify", reflect.TypeOf((*MockAuth)(nil).AsyncVerify), arg0)
}

// CanDeduct mocks base method.
func (m *MockAuth) CanDeduct(arg0 context.Context, arg1 Database, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanDeduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanDeduct indicates an expected call of CanDeduct.
func (mr *MockAuthMockRecorder) CanDeduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanDeduct", reflect.TypeOf((*MockAuth)(nil).CanDeduct), arg0, arg1, arg2)
}

// Deduct mocks base method.
func (m *MockAuth) Deduct(arg0 context.Context, arg1 Database, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deduct indicates an expected call of Deduct.
func (mr *MockAuthMockRecorder) Deduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deduct", reflect.TypeOf((*MockAuth)(nil).Deduct), arg0, arg1, arg2)
}

// Marshal mocks base method.
func (m *MockAuth) Marshal(arg0 *codec.Packer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Marshal", arg0)
}

// Marshal indicates an expected call of Marshal.
func (mr *MockAuthMockRecorder) Marshal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockAuth)(nil).Marshal), arg0)
}

// MaxUnits mocks base method.
func (m *MockAuth) MaxUnits(arg0 Rules) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxUnits", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// MaxUnits indicates an expected call of MaxUnits.
func (mr *MockAuthMockRecorder) MaxUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxUnits", reflect.TypeOf((*MockAuth)(nil).MaxUnits), arg0)
}

// Payer mocks base method.
func (m *MockAuth) Payer() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payer")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Payer indicates an expected call of Payer.
func (mr *MockAuthMockRecorder) Payer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payer", reflect.TypeOf((*MockAuth)(nil).Payer))
}

// Refund mocks base method.
func (m *MockAuth) Refund(arg0 context.Context, arg1 Database, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refund", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refund indicates an expected call of Refund.
func (mr *MockAuthMockRecorder) Refund(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refund", reflect.TypeOf((*MockAuth)(nil).Refund), arg0, arg1, arg2)
}

// StateKeys mocks base method.
func (m *MockAuth) StateKeys() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateKeys")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// StateKeys indicates an expected call of StateKeys.
func (mr *MockAuthMockRecorder) StateKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateKeys", reflect.TypeOf((*MockAuth)(nil).StateKeys))
}

// ValidRange mocks base method.
func (m *MockAuth) ValidRange(arg0 Rules) (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidRange", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// ValidRange indicates an expected call of ValidRange.
func (mr *MockAuthMockRecorder) ValidRange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidRange", reflect.TypeOf((*MockAuth)(nil).ValidRange), arg0)
}

// Verify mocks base method.
func (m *MockAuth) Verify(arg0 context.Context, arg1 Rules, arg2 Database, arg3 Action) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockAuthMockRecorder) Verify(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuth)(nil).Verify), arg0, arg1, arg2, arg3)
}
