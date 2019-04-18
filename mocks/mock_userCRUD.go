// Code generated by MockGen. DO NOT EDIT.
// Source: team-project/database (interfaces: UserCRUD)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
	data "team-project/services/data"
)

// MockUserCRUD is a mock of UserCRUD interface
type MockUserCRUD struct {
	ctrl     *gomock.Controller
	recorder *MockUserCRUDMockRecorder
}

// MockUserCRUDMockRecorder is the mock recorder for MockUserCRUD
type MockUserCRUDMockRecorder struct {
	mock *MockUserCRUD
}

// NewMockUserCRUD creates a new mock instance
func NewMockUserCRUD(ctrl *gomock.Controller) *MockUserCRUD {
	mock := &MockUserCRUD{ctrl: ctrl}
	mock.recorder = &MockUserCRUDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserCRUD) EXPECT() *MockUserCRUDMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockUserCRUD) AddUser(arg0 data.User) (data.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0)
	ret0, _ := ret[0].(data.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser
func (mr *MockUserCRUDMockRecorder) AddUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserCRUD)(nil).AddUser), arg0)
}

// DeleteUser mocks base method
func (m *MockUserCRUD) DeleteUser(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserCRUDMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserCRUD)(nil).DeleteUser), arg0)
}

// GetAllUsers mocks base method
func (m *MockUserCRUD) GetAllUsers() ([]data.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]data.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUserCRUDMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserCRUD)(nil).GetAllUsers))
}

// GetUserPassword mocks base method
func (m *MockUserCRUD) GetUserPassword(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPassword indicates an expected call of GetUserPassword
func (mr *MockUserCRUDMockRecorder) GetUserPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPassword", reflect.TypeOf((*MockUserCRUD)(nil).GetUserPassword), arg0)
}

// GetUserRole mocks base method
func (m *MockUserCRUD) GetUserRole(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole
func (mr *MockUserCRUDMockRecorder) GetUserRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockUserCRUD)(nil).GetUserRole), arg0)
}

// UpdateUser mocks base method
func (m *MockUserCRUD) UpdateUser(arg0 data.User, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserCRUDMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserCRUD)(nil).UpdateUser), arg0, arg1)
}
