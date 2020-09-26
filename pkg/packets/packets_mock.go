// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/packets/packets.go

// Package packets is a generated GoMock package.
package packets

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockPacket is a mock of Packet interface
type MockPacket struct {
	ctrl     *gomock.Controller
	recorder *MockPacketMockRecorder
}

// MockPacketMockRecorder is the mock recorder for MockPacket
type MockPacketMockRecorder struct {
	mock *MockPacket
}

// NewMockPacket creates a new mock instance
func NewMockPacket(ctrl *gomock.Controller) *MockPacket {
	mock := &MockPacket{ctrl: ctrl}
	mock.recorder = &MockPacketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacket) EXPECT() *MockPacketMockRecorder {
	return m.recorder
}

// Pack mocks base method
func (m *MockPacket) Pack(w io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pack", w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pack indicates an expected call of Pack
func (mr *MockPacketMockRecorder) Pack(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pack", reflect.TypeOf((*MockPacket)(nil).Pack), w)
}

// Unpack mocks base method
func (m *MockPacket) Unpack(r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unpack indicates an expected call of Unpack
func (mr *MockPacketMockRecorder) Unpack(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockPacket)(nil).Unpack), r)
}

// String mocks base method
func (m *MockPacket) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockPacketMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockPacket)(nil).String))
}
