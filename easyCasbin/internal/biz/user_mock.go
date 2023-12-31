// Code generated by MockGen. DO NOT EDIT.
// Source: easyCasbin/internal/biz (interfaces: UserRepo)

// Package biz is a generated GoMock package.
package biz

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepo) CreateUser(arg0 context.Context, arg1 *User) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepo)(nil).CreateUser), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockUserRepo) GetUserById(arg0 context.Context, arg1 int64) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepoMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepo)(nil).GetUserById), arg0, arg1)
}

// GetUserByName mocks base method.
func (m *MockUserRepo) GetUserByName(arg0 context.Context, arg1 string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserRepoMockRecorder) GetUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserRepo)(nil).GetUserByName), arg0, arg1)
}

// ListUser mocks base method.
func (m *MockUserRepo) ListUser(arg0 context.Context, arg1, arg2 int) ([]*User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUser indicates an expected call of ListUser.
func (mr *MockUserRepoMockRecorder) ListUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockUserRepo)(nil).ListUser), arg0, arg1, arg2)
}

// UpdateUser mocks base method.
func (m *MockUserRepo) UpdateUser(arg0 context.Context, arg1 *User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepoMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepo)(nil).UpdateUser), arg0, arg1)
}

// UserByMobile mocks base method.
func (m *MockUserRepo) UserByMobile(arg0 context.Context, arg1 string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserByMobile", arg0, arg1)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserByMobile indicates an expected call of UserByMobile.
func (mr *MockUserRepoMockRecorder) UserByMobile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserByMobile", reflect.TypeOf((*MockUserRepo)(nil).UserByMobile), arg0, arg1)
}
