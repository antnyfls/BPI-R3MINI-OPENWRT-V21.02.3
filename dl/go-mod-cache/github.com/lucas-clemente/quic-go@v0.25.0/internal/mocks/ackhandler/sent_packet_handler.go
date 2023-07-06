// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go/internal/ackhandler (interfaces: SentPacketHandler)

// Package mockackhandler is a generated GoMock package.
package mockackhandler

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ackhandler "github.com/lucas-clemente/quic-go/internal/ackhandler"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
	wire "github.com/lucas-clemente/quic-go/internal/wire"
)

// MockSentPacketHandler is a mock of SentPacketHandler interface.
type MockSentPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSentPacketHandlerMockRecorder
}

// MockSentPacketHandlerMockRecorder is the mock recorder for MockSentPacketHandler.
type MockSentPacketHandlerMockRecorder struct {
	mock *MockSentPacketHandler
}

// NewMockSentPacketHandler creates a new mock instance.
func NewMockSentPacketHandler(ctrl *gomock.Controller) *MockSentPacketHandler {
	mock := &MockSentPacketHandler{ctrl: ctrl}
	mock.recorder = &MockSentPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentPacketHandler) EXPECT() *MockSentPacketHandlerMockRecorder {
	return m.recorder
}

// DropPackets mocks base method.
func (m *MockSentPacketHandler) DropPackets(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DropPackets", arg0)
}

// DropPackets indicates an expected call of DropPackets.
func (mr *MockSentPacketHandlerMockRecorder) DropPackets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropPackets", reflect.TypeOf((*MockSentPacketHandler)(nil).DropPackets), arg0)
}

// GetLossDetectionTimeout mocks base method.
func (m *MockSentPacketHandler) GetLossDetectionTimeout() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLossDetectionTimeout")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLossDetectionTimeout indicates an expected call of GetLossDetectionTimeout.
func (mr *MockSentPacketHandlerMockRecorder) GetLossDetectionTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLossDetectionTimeout", reflect.TypeOf((*MockSentPacketHandler)(nil).GetLossDetectionTimeout))
}

// HasPacingBudget mocks base method.
func (m *MockSentPacketHandler) HasPacingBudget() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPacingBudget")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPacingBudget indicates an expected call of HasPacingBudget.
func (mr *MockSentPacketHandlerMockRecorder) HasPacingBudget() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPacingBudget", reflect.TypeOf((*MockSentPacketHandler)(nil).HasPacingBudget))
}

// OnLossDetectionTimeout mocks base method.
func (m *MockSentPacketHandler) OnLossDetectionTimeout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnLossDetectionTimeout")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnLossDetectionTimeout indicates an expected call of OnLossDetectionTimeout.
func (mr *MockSentPacketHandlerMockRecorder) OnLossDetectionTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnLossDetectionTimeout", reflect.TypeOf((*MockSentPacketHandler)(nil).OnLossDetectionTimeout))
}

// PeekPacketNumber mocks base method.
func (m *MockSentPacketHandler) PeekPacketNumber(arg0 protocol.EncryptionLevel) (protocol.PacketNumber, protocol.PacketNumberLen) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekPacketNumber", arg0)
	ret0, _ := ret[0].(protocol.PacketNumber)
	ret1, _ := ret[1].(protocol.PacketNumberLen)
	return ret0, ret1
}

// PeekPacketNumber indicates an expected call of PeekPacketNumber.
func (mr *MockSentPacketHandlerMockRecorder) PeekPacketNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekPacketNumber", reflect.TypeOf((*MockSentPacketHandler)(nil).PeekPacketNumber), arg0)
}

// PopPacketNumber mocks base method.
func (m *MockSentPacketHandler) PopPacketNumber(arg0 protocol.EncryptionLevel) protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopPacketNumber", arg0)
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// PopPacketNumber indicates an expected call of PopPacketNumber.
func (mr *MockSentPacketHandlerMockRecorder) PopPacketNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopPacketNumber", reflect.TypeOf((*MockSentPacketHandler)(nil).PopPacketNumber), arg0)
}

// QueueProbePacket mocks base method.
func (m *MockSentPacketHandler) QueueProbePacket(arg0 protocol.EncryptionLevel) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueProbePacket", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// QueueProbePacket indicates an expected call of QueueProbePacket.
func (mr *MockSentPacketHandlerMockRecorder) QueueProbePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueProbePacket", reflect.TypeOf((*MockSentPacketHandler)(nil).QueueProbePacket), arg0)
}

// ReceivedAck mocks base method.
func (m *MockSentPacketHandler) ReceivedAck(arg0 *wire.AckFrame, arg1 protocol.EncryptionLevel, arg2 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedAck", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceivedAck indicates an expected call of ReceivedAck.
func (mr *MockSentPacketHandlerMockRecorder) ReceivedAck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedAck", reflect.TypeOf((*MockSentPacketHandler)(nil).ReceivedAck), arg0, arg1, arg2)
}

// ReceivedBytes mocks base method.
func (m *MockSentPacketHandler) ReceivedBytes(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedBytes", arg0)
}

// ReceivedBytes indicates an expected call of ReceivedBytes.
func (mr *MockSentPacketHandlerMockRecorder) ReceivedBytes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedBytes", reflect.TypeOf((*MockSentPacketHandler)(nil).ReceivedBytes), arg0)
}

// ResetForRetry mocks base method.
func (m *MockSentPacketHandler) ResetForRetry() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetForRetry")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetForRetry indicates an expected call of ResetForRetry.
func (mr *MockSentPacketHandlerMockRecorder) ResetForRetry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetForRetry", reflect.TypeOf((*MockSentPacketHandler)(nil).ResetForRetry))
}

// SendMode mocks base method.
func (m *MockSentPacketHandler) SendMode() ackhandler.SendMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMode")
	ret0, _ := ret[0].(ackhandler.SendMode)
	return ret0
}

// SendMode indicates an expected call of SendMode.
func (mr *MockSentPacketHandlerMockRecorder) SendMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMode", reflect.TypeOf((*MockSentPacketHandler)(nil).SendMode))
}

// SentPacket mocks base method.
func (m *MockSentPacketHandler) SentPacket(arg0 *ackhandler.Packet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SentPacket", arg0)
}

// SentPacket indicates an expected call of SentPacket.
func (mr *MockSentPacketHandlerMockRecorder) SentPacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SentPacket", reflect.TypeOf((*MockSentPacketHandler)(nil).SentPacket), arg0)
}

// SetHandshakeConfirmed mocks base method.
func (m *MockSentPacketHandler) SetHandshakeConfirmed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandshakeConfirmed")
}

// SetHandshakeConfirmed indicates an expected call of SetHandshakeConfirmed.
func (mr *MockSentPacketHandlerMockRecorder) SetHandshakeConfirmed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandshakeConfirmed", reflect.TypeOf((*MockSentPacketHandler)(nil).SetHandshakeConfirmed))
}

// SetMaxDatagramSize mocks base method.
func (m *MockSentPacketHandler) SetMaxDatagramSize(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxDatagramSize", arg0)
}

// SetMaxDatagramSize indicates an expected call of SetMaxDatagramSize.
func (mr *MockSentPacketHandlerMockRecorder) SetMaxDatagramSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxDatagramSize", reflect.TypeOf((*MockSentPacketHandler)(nil).SetMaxDatagramSize), arg0)
}

// TimeUntilSend mocks base method.
func (m *MockSentPacketHandler) TimeUntilSend() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeUntilSend")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// TimeUntilSend indicates an expected call of TimeUntilSend.
func (mr *MockSentPacketHandlerMockRecorder) TimeUntilSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeUntilSend", reflect.TypeOf((*MockSentPacketHandler)(nil).TimeUntilSend))
}