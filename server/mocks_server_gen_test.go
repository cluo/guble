// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/smancke/guble/server (interfaces: PubSubSource,MessageSink,WSConn,Startable,Stopable,SetRouter,SetMessageEntry,Endpoint,SetKVStore,SetMessageStore)

package server

import (
	gomock "github.com/golang/mock/gomock"
	guble "github.com/smancke/guble/guble"
	
	store "github.com/smancke/guble/store"
	http "net/http"
)

// Mock of PubSubSource interface
type MockPubSubSource struct {
	ctrl     *gomock.Controller
	recorder *_MockPubSubSourceRecorder
}

// Recorder for MockPubSubSource (not exported)
type _MockPubSubSourceRecorder struct {
	mock *MockPubSubSource
}

func NewMockPubSubSource(ctrl *gomock.Controller) *MockPubSubSource {
	mock := &MockPubSubSource{ctrl: ctrl}
	mock.recorder = &_MockPubSubSourceRecorder{mock}
	return mock
}

func (_m *MockPubSubSource) EXPECT() *_MockPubSubSourceRecorder {
	return _m.recorder
}

func (_m *MockPubSubSource) Subscribe(_param0 *Route) (*Route, error) {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(*Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPubSubSourceRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Subscribe", arg0)
}

func (_m *MockPubSubSource) Unsubscribe(_param0 *Route) {
	_m.ctrl.Call(_m, "Unsubscribe", _param0)
}

func (_mr *_MockPubSubSourceRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unsubscribe", arg0)
}

// Mock of MessageSink interface
type MockMessageSink struct {
	ctrl     *gomock.Controller
	recorder *_MockMessageSinkRecorder
}

// Recorder for MockMessageSink (not exported)
type _MockMessageSinkRecorder struct {
	mock *MockMessageSink
}

func NewMockMessageSink(ctrl *gomock.Controller) *MockMessageSink {
	mock := &MockMessageSink{ctrl: ctrl}
	mock.recorder = &_MockMessageSinkRecorder{mock}
	return mock
}

func (_m *MockMessageSink) EXPECT() *_MockMessageSinkRecorder {
	return _m.recorder
}

func (_m *MockMessageSink) HandleMessage(_param0 *guble.Message) error {
	ret := _m.ctrl.Call(_m, "HandleMessage", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockMessageSinkRecorder) HandleMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleMessage", arg0)
}

// Mock of WSConn interface
type MockWSConn struct {
	ctrl     *gomock.Controller
	recorder *_MockWSConnRecorder
}

// Recorder for MockWSConn (not exported)
type _MockWSConnRecorder struct {
	mock *MockWSConn
}

func NewMockWSConn(ctrl *gomock.Controller) *MockWSConn {
	mock := &MockWSConn{ctrl: ctrl}
	mock.recorder = &_MockWSConnRecorder{mock}
	return mock
}

func (_m *MockWSConn) EXPECT() *_MockWSConnRecorder {
	return _m.recorder
}

func (_m *MockWSConn) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockWSConnRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockWSConn) Receive(_param0 *[]byte) error {
	ret := _m.ctrl.Call(_m, "Receive", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWSConnRecorder) Receive(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Receive", arg0)
}

func (_m *MockWSConn) Send(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWSConnRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0)
}

// Mock of Startable interface
type MockStartable struct {
	ctrl     *gomock.Controller
	recorder *_MockStartableRecorder
}

// Recorder for MockStartable (not exported)
type _MockStartableRecorder struct {
	mock *MockStartable
}

func NewMockStartable(ctrl *gomock.Controller) *MockStartable {
	mock := &MockStartable{ctrl: ctrl}
	mock.recorder = &_MockStartableRecorder{mock}
	return mock
}

func (_m *MockStartable) EXPECT() *_MockStartableRecorder {
	return _m.recorder
}

func (_m *MockStartable) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStartableRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

// Mock of Stopable interface
type MockStopable struct {
	ctrl     *gomock.Controller
	recorder *_MockStopableRecorder
}

// Recorder for MockStopable (not exported)
type _MockStopableRecorder struct {
	mock *MockStopable
}

func NewMockStopable(ctrl *gomock.Controller) *MockStopable {
	mock := &MockStopable{ctrl: ctrl}
	mock.recorder = &_MockStopableRecorder{mock}
	return mock
}

func (_m *MockStopable) EXPECT() *_MockStopableRecorder {
	return _m.recorder
}

func (_m *MockStopable) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStopableRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

// Mock of SetRouter interface
type MockSetRouter struct {
	ctrl     *gomock.Controller
	recorder *_MockSetRouterRecorder
}

// Recorder for MockSetRouter (not exported)
type _MockSetRouterRecorder struct {
	mock *MockSetRouter
}

func NewMockSetRouter(ctrl *gomock.Controller) *MockSetRouter {
	mock := &MockSetRouter{ctrl: ctrl}
	mock.recorder = &_MockSetRouterRecorder{mock}
	return mock
}

func (_m *MockSetRouter) EXPECT() *_MockSetRouterRecorder {
	return _m.recorder
}

func (_m *MockSetRouter) SetRouter(_param0 PubSubSource) {
	_m.ctrl.Call(_m, "SetRouter", _param0)
}

func (_mr *_MockSetRouterRecorder) SetRouter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRouter", arg0)
}

// Mock of SetMessageEntry interface
type MockSetMessageEntry struct {
	ctrl     *gomock.Controller
	recorder *_MockSetMessageEntryRecorder
}

// Recorder for MockSetMessageEntry (not exported)
type _MockSetMessageEntryRecorder struct {
	mock *MockSetMessageEntry
}

func NewMockSetMessageEntry(ctrl *gomock.Controller) *MockSetMessageEntry {
	mock := &MockSetMessageEntry{ctrl: ctrl}
	mock.recorder = &_MockSetMessageEntryRecorder{mock}
	return mock
}

func (_m *MockSetMessageEntry) EXPECT() *_MockSetMessageEntryRecorder {
	return _m.recorder
}

func (_m *MockSetMessageEntry) SetMessageEntry(_param0 MessageSink) {
	_m.ctrl.Call(_m, "SetMessageEntry", _param0)
}

func (_mr *_MockSetMessageEntryRecorder) SetMessageEntry(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMessageEntry", arg0)
}

// Mock of Endpoint interface
type MockEndpoint struct {
	ctrl     *gomock.Controller
	recorder *_MockEndpointRecorder
}

// Recorder for MockEndpoint (not exported)
type _MockEndpointRecorder struct {
	mock *MockEndpoint
}

func NewMockEndpoint(ctrl *gomock.Controller) *MockEndpoint {
	mock := &MockEndpoint{ctrl: ctrl}
	mock.recorder = &_MockEndpointRecorder{mock}
	return mock
}

func (_m *MockEndpoint) EXPECT() *_MockEndpointRecorder {
	return _m.recorder
}

func (_m *MockEndpoint) GetPrefix() string {
	ret := _m.ctrl.Call(_m, "GetPrefix")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockEndpointRecorder) GetPrefix() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPrefix")
}

func (_m *MockEndpoint) ServeHTTP(_param0 http.ResponseWriter, _param1 *http.Request) {
	_m.ctrl.Call(_m, "ServeHTTP", _param0, _param1)
}

func (_mr *_MockEndpointRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServeHTTP", arg0, arg1)
}

// Mock of SetKVStore interface
type MockSetKVStore struct {
	ctrl     *gomock.Controller
	recorder *_MockSetKVStoreRecorder
}

// Recorder for MockSetKVStore (not exported)
type _MockSetKVStoreRecorder struct {
	mock *MockSetKVStore
}

func NewMockSetKVStore(ctrl *gomock.Controller) *MockSetKVStore {
	mock := &MockSetKVStore{ctrl: ctrl}
	mock.recorder = &_MockSetKVStoreRecorder{mock}
	return mock
}

func (_m *MockSetKVStore) EXPECT() *_MockSetKVStoreRecorder {
	return _m.recorder
}

func (_m *MockSetKVStore) SetKVStore(_param0 store.KVStore) {
	_m.ctrl.Call(_m, "SetKVStore", _param0)
}

func (_mr *_MockSetKVStoreRecorder) SetKVStore(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetKVStore", arg0)
}

// Mock of SetMessageStore interface
type MockSetMessageStore struct {
	ctrl     *gomock.Controller
	recorder *_MockSetMessageStoreRecorder
}

// Recorder for MockSetMessageStore (not exported)
type _MockSetMessageStoreRecorder struct {
	mock *MockSetMessageStore
}

func NewMockSetMessageStore(ctrl *gomock.Controller) *MockSetMessageStore {
	mock := &MockSetMessageStore{ctrl: ctrl}
	mock.recorder = &_MockSetMessageStoreRecorder{mock}
	return mock
}

func (_m *MockSetMessageStore) EXPECT() *_MockSetMessageStoreRecorder {
	return _m.recorder
}

func (_m *MockSetMessageStore) SetMessageStore(_param0 store.MessageStore) {
	_m.ctrl.Call(_m, "SetMessageStore", _param0)
}

func (_mr *_MockSetMessageStoreRecorder) SetMessageStore(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetMessageStore", arg0)
}
