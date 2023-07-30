// Code generated by MockGen. DO NOT EDIT.
// Source: internal/server/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	websocket "nhooyr.io/websocket"
)

// MockSubscriber is a mock of Subscriber interface.
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber.
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance.
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSubscriber) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSubscriberMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSubscriber)(nil).Close))
}

// Listen mocks base method.
func (m *MockSubscriber) Listen() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen")
	ret0, _ := ret[0].(error)
	return ret0
}

// Listen indicates an expected call of Listen.
func (mr *MockSubscriberMockRecorder) Listen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockSubscriber)(nil).Listen))
}

// MockSocketPublisher is a mock of SocketPublisher interface.
type MockSocketPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockSocketPublisherMockRecorder
}

// MockSocketPublisherMockRecorder is the mock recorder for MockSocketPublisher.
type MockSocketPublisherMockRecorder struct {
	mock *MockSocketPublisher
}

// NewMockSocketPublisher creates a new mock instance.
func NewMockSocketPublisher(ctrl *gomock.Controller) *MockSocketPublisher {
	mock := &MockSocketPublisher{ctrl: ctrl}
	mock.recorder = &MockSocketPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocketPublisher) EXPECT() *MockSocketPublisherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSocketPublisher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSocketPublisherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSocketPublisher)(nil).Close))
}

// PublishLoop mocks base method.
func (m *MockSocketPublisher) PublishLoop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishLoop")
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishLoop indicates an expected call of PublishLoop.
func (mr *MockSocketPublisherMockRecorder) PublishLoop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishLoop", reflect.TypeOf((*MockSocketPublisher)(nil).PublishLoop))
}

// SendMessage mocks base method.
func (m *MockSocketPublisher) SendMessage(msg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockSocketPublisherMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSocketPublisher)(nil).SendMessage), msg)
}

// MockWebsocketConn is a mock of WebsocketConn interface.
type MockWebsocketConn struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketConnMockRecorder
}

// MockWebsocketConnMockRecorder is the mock recorder for MockWebsocketConn.
type MockWebsocketConnMockRecorder struct {
	mock *MockWebsocketConn
}

// NewMockWebsocketConn creates a new mock instance.
func NewMockWebsocketConn(ctrl *gomock.Controller) *MockWebsocketConn {
	mock := &MockWebsocketConn{ctrl: ctrl}
	mock.recorder = &MockWebsocketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebsocketConn) EXPECT() *MockWebsocketConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockWebsocketConn) Close(code websocket.StatusCode, reason string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", code, reason)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockWebsocketConnMockRecorder) Close(code, reason interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWebsocketConn)(nil).Close), code, reason)
}

// CloseRead mocks base method.
func (m *MockWebsocketConn) CloseRead(ctx context.Context) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseRead", ctx)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// CloseRead indicates an expected call of CloseRead.
func (mr *MockWebsocketConnMockRecorder) CloseRead(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseRead", reflect.TypeOf((*MockWebsocketConn)(nil).CloseRead), ctx)
}

// Ping mocks base method.
func (m *MockWebsocketConn) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockWebsocketConnMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockWebsocketConn)(nil).Ping), ctx)
}

// Read mocks base method.
func (m *MockWebsocketConn) Read(ctx context.Context) (websocket.MessageType, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(websocket.MessageType)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Read indicates an expected call of Read.
func (mr *MockWebsocketConnMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockWebsocketConn)(nil).Read), ctx)
}

// Reader mocks base method.
func (m *MockWebsocketConn) Reader(ctx context.Context) (websocket.MessageType, io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", ctx)
	ret0, _ := ret[0].(websocket.MessageType)
	ret1, _ := ret[1].(io.Reader)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Reader indicates an expected call of Reader.
func (mr *MockWebsocketConnMockRecorder) Reader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockWebsocketConn)(nil).Reader), ctx)
}

// SetReadLimit mocks base method.
func (m *MockWebsocketConn) SetReadLimit(n int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadLimit", n)
}

// SetReadLimit indicates an expected call of SetReadLimit.
func (mr *MockWebsocketConnMockRecorder) SetReadLimit(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLimit", reflect.TypeOf((*MockWebsocketConn)(nil).SetReadLimit), n)
}

// Subprotocol mocks base method.
func (m *MockWebsocketConn) Subprotocol() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subprotocol")
	ret0, _ := ret[0].(string)
	return ret0
}

// Subprotocol indicates an expected call of Subprotocol.
func (mr *MockWebsocketConnMockRecorder) Subprotocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subprotocol", reflect.TypeOf((*MockWebsocketConn)(nil).Subprotocol))
}

// Write mocks base method.
func (m *MockWebsocketConn) Write(ctx context.Context, typ websocket.MessageType, p []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, typ, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockWebsocketConnMockRecorder) Write(ctx, typ, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWebsocketConn)(nil).Write), ctx, typ, p)
}

// Writer mocks base method.
func (m *MockWebsocketConn) Writer(ctx context.Context, typ websocket.MessageType) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer", ctx, typ)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Writer indicates an expected call of Writer.
func (mr *MockWebsocketConnMockRecorder) Writer(ctx, typ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockWebsocketConn)(nil).Writer), ctx, typ)
}

// MockRedisPubSubConn is a mock of RedisPubSubConn interface.
type MockRedisPubSubConn struct {
	ctrl     *gomock.Controller
	recorder *MockRedisPubSubConnMockRecorder
}

// MockRedisPubSubConnMockRecorder is the mock recorder for MockRedisPubSubConn.
type MockRedisPubSubConnMockRecorder struct {
	mock *MockRedisPubSubConn
}

// NewMockRedisPubSubConn creates a new mock instance.
func NewMockRedisPubSubConn(ctrl *gomock.Controller) *MockRedisPubSubConn {
	mock := &MockRedisPubSubConn{ctrl: ctrl}
	mock.recorder = &MockRedisPubSubConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisPubSubConn) EXPECT() *MockRedisPubSubConnMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisPubSubConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisPubSubConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisPubSubConn)(nil).Close))
}

// PSubscribe mocks base method.
func (m *MockRedisPubSubConn) PSubscribe(channel ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range channel {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PSubscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PSubscribe indicates an expected call of PSubscribe.
func (mr *MockRedisPubSubConnMockRecorder) PSubscribe(channel ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PSubscribe", reflect.TypeOf((*MockRedisPubSubConn)(nil).PSubscribe), channel...)
}

// PUnsubscribe mocks base method.
func (m *MockRedisPubSubConn) PUnsubscribe(channel ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range channel {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUnsubscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PUnsubscribe indicates an expected call of PUnsubscribe.
func (mr *MockRedisPubSubConnMockRecorder) PUnsubscribe(channel ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUnsubscribe", reflect.TypeOf((*MockRedisPubSubConn)(nil).PUnsubscribe), channel...)
}

// Ping mocks base method.
func (m *MockRedisPubSubConn) Ping(data string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRedisPubSubConnMockRecorder) Ping(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRedisPubSubConn)(nil).Ping), data)
}

// Receive mocks base method.
func (m *MockRedisPubSubConn) Receive() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Receive indicates an expected call of Receive.
func (mr *MockRedisPubSubConnMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockRedisPubSubConn)(nil).Receive))
}

// ReceiveContext mocks base method.
func (m *MockRedisPubSubConn) ReceiveContext(ctx context.Context) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveContext", ctx)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// ReceiveContext indicates an expected call of ReceiveContext.
func (mr *MockRedisPubSubConnMockRecorder) ReceiveContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveContext", reflect.TypeOf((*MockRedisPubSubConn)(nil).ReceiveContext), ctx)
}

// ReceiveWithTimeout mocks base method.
func (m *MockRedisPubSubConn) ReceiveWithTimeout(timeout time.Duration) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveWithTimeout", timeout)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// ReceiveWithTimeout indicates an expected call of ReceiveWithTimeout.
func (mr *MockRedisPubSubConnMockRecorder) ReceiveWithTimeout(timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveWithTimeout", reflect.TypeOf((*MockRedisPubSubConn)(nil).ReceiveWithTimeout), timeout)
}

// Subscribe mocks base method.
func (m *MockRedisPubSubConn) Subscribe(channel ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range channel {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockRedisPubSubConnMockRecorder) Subscribe(channel ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockRedisPubSubConn)(nil).Subscribe), channel...)
}

// Unsubscribe mocks base method.
func (m *MockRedisPubSubConn) Unsubscribe(channel ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range channel {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unsubscribe", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockRedisPubSubConnMockRecorder) Unsubscribe(channel ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockRedisPubSubConn)(nil).Unsubscribe), channel...)
}