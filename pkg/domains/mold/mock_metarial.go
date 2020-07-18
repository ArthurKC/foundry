// Code generated by MockGen. DO NOT EDIT.
// Source: metarial.go

// Package mold is a generated GoMock package.
package mold

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMaterial is a mock of Material interface
type MockMaterial struct {
	ctrl     *gomock.Controller
	recorder *MockMaterialMockRecorder
}

// MockMaterialMockRecorder is the mock recorder for MockMaterial
type MockMaterialMockRecorder struct {
	mock *MockMaterial
}

// NewMockMaterial creates a new mock instance
func NewMockMaterial(ctrl *gomock.Controller) *MockMaterial {
	mock := &MockMaterial{ctrl: ctrl}
	mock.recorder = &MockMaterialMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaterial) EXPECT() *MockMaterialMockRecorder {
	return m.recorder
}

// Parameters mocks base method
func (m *MockMaterial) Parameters() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Parameters indicates an expected call of Parameters
func (mr *MockMaterialMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockMaterial)(nil).Parameters))
}