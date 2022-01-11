// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/exec/types.go

// Package exec is a generated GoMock package.
package exec

import (
	reflect "reflect"

	dns "github.com/alibaba/kt-connect/pkg/kt/exec/dns"
	sshchannel "github.com/alibaba/kt-connect/pkg/kt/exec/sshchannel"
	sshuttle "github.com/alibaba/kt-connect/pkg/kt/exec/sshuttle"
	tun "github.com/alibaba/kt-connect/pkg/kt/exec/tun"
	gomock "github.com/golang/mock/gomock"
)

// MockCliInterface is a mock of CliInterface interface.
type MockCliInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCliInterfaceMockRecorder
}

// MockCliInterfaceMockRecorder is the mock recorder for MockCliInterface.
type MockCliInterfaceMockRecorder struct {
	mock *MockCliInterface
}

// NewMockCliInterface creates a new mock instance.
func NewMockCliInterface(ctrl *gomock.Controller) *MockCliInterface {
	mock := &MockCliInterface{ctrl: ctrl}
	mock.recorder = &MockCliInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCliInterface) EXPECT() *MockCliInterfaceMockRecorder {
	return m.recorder
}

// DnsConfig mocks base method.
func (m *MockCliInterface) DnsConfig() dns.DnsConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DnsConfig")
	ret0, _ := ret[0].(dns.DnsConfig)
	return ret0
}

// DnsConfig indicates an expected call of DnsConfig.
func (mr *MockCliInterfaceMockRecorder) DnsConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DnsConfig", reflect.TypeOf((*MockCliInterface)(nil).DnsConfig))
}

// SshChannel mocks base method.
func (m *MockCliInterface) SshChannel() sshchannel.Channel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshChannel")
	ret0, _ := ret[0].(sshchannel.Channel)
	return ret0
}

// SshChannel indicates an expected call of SshChannel.
func (mr *MockCliInterfaceMockRecorder) SshChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshChannel", reflect.TypeOf((*MockCliInterface)(nil).SshChannel))
}

// Sshuttle mocks base method.
func (m *MockCliInterface) Sshuttle() sshuttle.Sshuttle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sshuttle")
	ret0, _ := ret[0].(sshuttle.Sshuttle)
	return ret0
}

// Sshuttle indicates an expected call of Sshuttle.
func (mr *MockCliInterfaceMockRecorder) Sshuttle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sshuttle", reflect.TypeOf((*MockCliInterface)(nil).Sshuttle))
}

// Tunnel mocks base method.
func (m *MockCliInterface) Tunnel() tun.Tunnel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tunnel")
	ret0, _ := ret[0].(tun.Tunnel)
	return ret0
}

// Tunnel indicates an expected call of Tunnel.
func (mr *MockCliInterfaceMockRecorder) Tunnel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tunnel", reflect.TypeOf((*MockCliInterface)(nil).Tunnel))
}
