// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/query/querier.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	query "github.com/tensorchord/envd-server/pkg/query"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateImageInfo mocks base method.
func (m *MockQuerier) CreateImageInfo(ctx context.Context, arg query.CreateImageInfoParams) (query.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImageInfo", ctx, arg)
	ret0, _ := ret[0].(query.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateImageInfo indicates an expected call of CreateImageInfo.
func (mr *MockQuerierMockRecorder) CreateImageInfo(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImageInfo", reflect.TypeOf((*MockQuerier)(nil).CreateImageInfo), ctx, arg)
}

// CreateKey mocks base method.
func (m *MockQuerier) CreateKey(ctx context.Context, arg query.CreateKeyParams) (query.CreateKeyRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", ctx, arg)
	ret0, _ := ret[0].(query.CreateKeyRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockQuerierMockRecorder) CreateKey(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockQuerier)(nil).CreateKey), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, arg query.CreateUserParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, arg)
}

// DeleteUser mocks base method.
func (m *MockQuerier) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockQuerierMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockQuerier)(nil).DeleteUser), ctx, id)
}

// GetImageInfoByDigest mocks base method.
func (m *MockQuerier) GetImageInfoByDigest(ctx context.Context, arg query.GetImageInfoByDigestParams) (query.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageInfoByDigest", ctx, arg)
	ret0, _ := ret[0].(query.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageInfoByDigest indicates an expected call of GetImageInfoByDigest.
func (mr *MockQuerierMockRecorder) GetImageInfoByDigest(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageInfoByDigest", reflect.TypeOf((*MockQuerier)(nil).GetImageInfoByDigest), ctx, arg)
}

// GetImageInfoByName mocks base method.
func (m *MockQuerier) GetImageInfoByName(ctx context.Context, arg query.GetImageInfoByNameParams) (query.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageInfoByName", ctx, arg)
	ret0, _ := ret[0].(query.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageInfoByName indicates an expected call of GetImageInfoByName.
func (mr *MockQuerierMockRecorder) GetImageInfoByName(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageInfoByName", reflect.TypeOf((*MockQuerier)(nil).GetImageInfoByName), ctx, arg)
}

// GetKey mocks base method.
func (m *MockQuerier) GetKey(ctx context.Context, arg query.GetKeyParams) (query.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, arg)
	ret0, _ := ret[0].(query.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey.
func (mr *MockQuerierMockRecorder) GetKey(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockQuerier)(nil).GetKey), ctx, arg)
}

// GetUser mocks base method.
func (m *MockQuerier) GetUser(ctx context.Context, loginName string) (query.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, loginName)
	ret0, _ := ret[0].(query.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQuerierMockRecorder) GetUser(ctx, loginName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQuerier)(nil).GetUser), ctx, loginName)
}

// ListImageByOwner mocks base method.
func (m *MockQuerier) ListImageByOwner(ctx context.Context, loginName string) ([]query.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImageByOwner", ctx, loginName)
	ret0, _ := ret[0].([]query.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImageByOwner indicates an expected call of ListImageByOwner.
func (mr *MockQuerierMockRecorder) ListImageByOwner(ctx, loginName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImageByOwner", reflect.TypeOf((*MockQuerier)(nil).ListImageByOwner), ctx, loginName)
}

// ListKeys mocks base method.
func (m *MockQuerier) ListKeys(ctx context.Context, loginName string) ([]query.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", ctx, loginName)
	ret0, _ := ret[0].([]query.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockQuerierMockRecorder) ListKeys(ctx, loginName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockQuerier)(nil).ListKeys), ctx, loginName)
}

// ListUsers mocks base method.
func (m *MockQuerier) ListUsers(ctx context.Context) ([]query.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx)
	ret0, _ := ret[0].([]query.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockQuerierMockRecorder) ListUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockQuerier)(nil).ListUsers), ctx)
}
