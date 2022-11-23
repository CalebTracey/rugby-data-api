// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-data-api/internal/dao/database/sixnations (interfaces: MapperI)

// Package sixnations is a generated GoMock package.
package sixnations

import (
	sql "database/sql"
	reflect "reflect"

	models "github.com/calebtracey/rugby-data-api/external/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMapperI is a mock of MapperI interface.
type MockMapperI struct {
	ctrl     *gomock.Controller
	recorder *MockMapperIMockRecorder
}

// MockMapperIMockRecorder is the mock recorder for MockMapperI.
type MockMapperIMockRecorder struct {
	mock *MockMapperI
}

// NewMockMapperI creates a new mock instance.
func NewMockMapperI(ctrl *gomock.Controller) *MockMapperI {
	mock := &MockMapperI{ctrl: ctrl}
	mock.recorder = &MockMapperIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapperI) EXPECT() *MockMapperIMockRecorder {
	return m.recorder
}

// CreatePSQLTeamsQuery mocks base method.
func (m *MockMapperI) CreatePSQLTeamsQuery(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePSQLTeamsQuery", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreatePSQLTeamsQuery indicates an expected call of CreatePSQLTeamsQuery.
func (mr *MockMapperIMockRecorder) CreatePSQLTeamsQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePSQLTeamsQuery", reflect.TypeOf((*MockMapperI)(nil).CreatePSQLTeamsQuery), arg0)
}

// MapPSQLRowsToTeamData mocks base method.
func (m *MockMapperI) MapPSQLRowsToTeamData(arg0 *sql.Rows) []models.TeamData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapPSQLRowsToTeamData", arg0)
	ret0, _ := ret[0].([]models.TeamData)
	return ret0
}

// MapPSQLRowsToTeamData indicates an expected call of MapPSQLRowsToTeamData.
func (mr *MockMapperIMockRecorder) MapPSQLRowsToTeamData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapPSQLRowsToTeamData", reflect.TypeOf((*MockMapperI)(nil).MapPSQLRowsToTeamData), arg0)
}
