// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/minhdanh/vault-template/api (interfaces: VaultClient)

// Package api is a generated GoMock package.
package api

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVaultClient is a mock of VaultClient interface
type MockVaultClient struct {
	ctrl     *gomock.Controller
	recorder *MockVaultClientMockRecorder
}

// MockVaultClientMockRecorder is the mock recorder for MockVaultClient
type MockVaultClientMockRecorder struct {
	mock *MockVaultClient
}

// NewMockVaultClient creates a new mock instance
func NewMockVaultClient(ctrl *gomock.Controller) *MockVaultClient {
	mock := &MockVaultClient{ctrl: ctrl}
	mock.recorder = &MockVaultClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVaultClient) EXPECT() *MockVaultClientMockRecorder {
	return m.recorder
}

// QuerySecretMap mocks base method
func (m *MockVaultClient) QuerySecretMap(arg0 string, arg1 ...string) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySecretMap", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySecret mocks base method
func (m *MockVaultClient) QuerySecret(arg0, arg1, string,arg2 ...string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySecret", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySecret indicates an expected call of QuerySecret
func (mr *MockVaultClientMockRecorder) QuerySecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySecret", reflect.TypeOf((*MockVaultClient)(nil).QuerySecret), arg0, arg1)
}

// QuerySecretMap indicates an expected call of QuerySecretMap
func (mr *MockVaultClientMockRecorder) QuerySecretMap(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySecretMap", reflect.TypeOf((*MockVaultClient)(nil).QuerySecretMap), arg0)
}
