// Code generated by MockGen. DO NOT EDIT.
// Source: ShamanEstBanan-GophKeeper-server/internal/service (interfaces: Storage)

// Package mock is a generated GoMock package.
package mock

import (
	entity "ShamanEstBanan-GophKeeper-server/internal/domain/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockStorage) AuthenticateUser(arg0 context.Context, arg1 entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockStorageMockRecorder) AuthenticateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockStorage)(nil).AuthenticateUser), arg0, arg1)
}

// CreateRecord mocks base method.
func (m *MockStorage) CreateRecord(arg0 context.Context, arg1 entity.Record) (*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", arg0, arg1)
	ret0, _ := ret[0].(*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockStorageMockRecorder) CreateRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockStorage)(nil).CreateRecord), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStorage) CreateUser(arg0 context.Context, arg1 entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorageMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser), arg0, arg1)
}

// DeleteRecord mocks base method.
func (m *MockStorage) DeleteRecord(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockStorageMockRecorder) DeleteRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockStorage)(nil).DeleteRecord), arg0, arg1, arg2)
}

// GetAllRecords mocks base method.
func (m *MockStorage) GetAllRecords(arg0 context.Context, arg1 string) (*[]entity.RecordInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecords", arg0, arg1)
	ret0, _ := ret[0].(*[]entity.RecordInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRecords indicates an expected call of GetAllRecords.
func (mr *MockStorageMockRecorder) GetAllRecords(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRecords", reflect.TypeOf((*MockStorage)(nil).GetAllRecords), arg0, arg1)
}

// GetRecord mocks base method.
func (m *MockStorage) GetRecord(arg0 context.Context, arg1, arg2 string) (*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord.
func (mr *MockStorageMockRecorder) GetRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockStorage)(nil).GetRecord), arg0, arg1, arg2)
}

// GetRecordsByType mocks base method.
func (m *MockStorage) GetRecordsByType(arg0 context.Context, arg1, arg2 string) (*[]entity.RecordInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsByType", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]entity.RecordInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsByType indicates an expected call of GetRecordsByType.
func (mr *MockStorageMockRecorder) GetRecordsByType(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsByType", reflect.TypeOf((*MockStorage)(nil).GetRecordsByType), arg0, arg1, arg2)
}

// UpdateRecord mocks base method.
func (m *MockStorage) UpdateRecord(arg0 context.Context, arg1 entity.Record) (*entity.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", arg0, arg1)
	ret0, _ := ret[0].(*entity.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRecord indicates an expected call of UpdateRecord.
func (mr *MockStorageMockRecorder) UpdateRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockStorage)(nil).UpdateRecord), arg0, arg1)
}
