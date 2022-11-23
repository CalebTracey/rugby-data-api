// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-data-api/internal/dao/database/sixnations (interfaces: SNDAOI)

// Package sixnations is a generated GoMock package.
package sixnations

import (
	context "context"
	reflect "reflect"

	response "github.com/calebtracey/rugby-data-api/external/models/response"
	gomock "github.com/golang/mock/gomock"
)

// MockSNDAOI is a mock of SNDAOI interface.
type MockSNDAOI struct {
	ctrl     *gomock.Controller
	recorder *MockSNDAOIMockRecorder
}

// MockSNDAOIMockRecorder is the mock recorder for MockSNDAOI.
type MockSNDAOIMockRecorder struct {
	mock *MockSNDAOI
}

// NewMockSNDAOI creates a new mock instance.
func NewMockSNDAOI(ctrl *gomock.Controller) *MockSNDAOI {
	mock := &MockSNDAOI{ctrl: ctrl}
	mock.recorder = &MockSNDAOIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSNDAOI) EXPECT() *MockSNDAOIMockRecorder {
	return m.recorder
}

// GetTeams mocks base method.
func (m *MockSNDAOI) GetTeams(arg0 context.Context, arg1 string) (response.CompetitionResponse, *response.ErrorLog) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeams", arg0, arg1)
	ret0, _ := ret[0].(response.CompetitionResponse)
	ret1, _ := ret[1].(*response.ErrorLog)
	return ret0, ret1
}

// GetTeams indicates an expected call of GetTeams.
func (mr *MockSNDAOIMockRecorder) GetTeams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeams", reflect.TypeOf((*MockSNDAOI)(nil).GetTeams), arg0, arg1)
}
