// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package database_mocks is a generated GoMock package.
package database_mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateExample mocks base method.
func (m *MockDatabase) CreateExample(ctx context.Context, description string) (*v1.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExample", ctx, description)
	ret0, _ := ret[0].(*v1.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExample indicates an expected call of CreateExample.
func (mr *MockDatabaseMockRecorder) CreateExample(ctx, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExample", reflect.TypeOf((*MockDatabase)(nil).CreateExample), ctx, description)
}

// FetchExamples mocks base method.
func (m *MockDatabase) FetchExamples(ctx context.Context) ([]*v1.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchExamples", ctx)
	ret0, _ := ret[0].([]*v1.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchExamples indicates an expected call of FetchExamples.
func (mr *MockDatabaseMockRecorder) FetchExamples(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchExamples", reflect.TypeOf((*MockDatabase)(nil).FetchExamples), ctx)
}

// GetExampleByID mocks base method.
func (m *MockDatabase) GetExampleByID(ctx context.Context, id string) (*v1.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExampleByID", ctx, id)
	ret0, _ := ret[0].(*v1.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExampleByID indicates an expected call of GetExampleByID.
func (mr *MockDatabaseMockRecorder) GetExampleByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExampleByID", reflect.TypeOf((*MockDatabase)(nil).GetExampleByID), ctx, id)
}
