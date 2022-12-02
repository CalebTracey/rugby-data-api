// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-data-api/internal/dao/psql (interfaces: MapperI)

// Package dbmocks is a generated GoMock package.
package dbmocks

import (
	sql "database/sql"
	reflect "reflect"

	dtos "github.com/calebtracey/rugby-models/pkg/dtos"
	leaderboard "github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	response "github.com/calebtracey/rugby-models/pkg/dtos/response"
	models "github.com/calebtracey/rugby-models/pkg/models"
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

// CreatePSQLLeaderboardByIdQuery mocks base method.
func (m *MockMapperI) CreatePSQLLeaderboardByIdQuery(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePSQLLeaderboardByIdQuery", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreatePSQLLeaderboardByIdQuery indicates an expected call of CreatePSQLLeaderboardByIdQuery.
func (mr *MockMapperIMockRecorder) CreatePSQLLeaderboardByIdQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePSQLLeaderboardByIdQuery", reflect.TypeOf((*MockMapperI)(nil).CreatePSQLLeaderboardByIdQuery), arg0)
}

// MapPSQLAllLeaderboardDataToResponse mocks base method.
func (m *MockMapperI) MapPSQLAllLeaderboardDataToResponse(arg0 models.PSQLLeaderboardDataList) leaderboard.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapPSQLAllLeaderboardDataToResponse", arg0)
	ret0, _ := ret[0].(leaderboard.Response)
	return ret0
}

// MapPSQLAllLeaderboardDataToResponse indicates an expected call of MapPSQLAllLeaderboardDataToResponse.
func (mr *MockMapperIMockRecorder) MapPSQLAllLeaderboardDataToResponse(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapPSQLAllLeaderboardDataToResponse", reflect.TypeOf((*MockMapperI)(nil).MapPSQLAllLeaderboardDataToResponse), arg0)
}

// MapPSQLLeaderboardDataToResponse mocks base method.
func (m *MockMapperI) MapPSQLLeaderboardDataToResponse(arg0, arg1 string, arg2 models.PSQLLeaderboardDataList) dtos.CompetitionLeaderboardData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapPSQLLeaderboardDataToResponse", arg0, arg1, arg2)
	ret0, _ := ret[0].(dtos.CompetitionLeaderboardData)
	return ret0
}

// MapPSQLLeaderboardDataToResponse indicates an expected call of MapPSQLLeaderboardDataToResponse.
func (mr *MockMapperIMockRecorder) MapPSQLLeaderboardDataToResponse(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapPSQLLeaderboardDataToResponse", reflect.TypeOf((*MockMapperI)(nil).MapPSQLLeaderboardDataToResponse), arg0, arg1, arg2)
}

// MapPSQLRowsToLeaderboardData mocks base method.
func (m *MockMapperI) MapPSQLRowsToLeaderboardData(arg0 *sql.Rows) (models.PSQLLeaderboardDataList, *response.ErrorLog) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapPSQLRowsToLeaderboardData", arg0)
	ret0, _ := ret[0].(models.PSQLLeaderboardDataList)
	ret1, _ := ret[1].(*response.ErrorLog)
	return ret0, ret1
}

// MapPSQLRowsToLeaderboardData indicates an expected call of MapPSQLRowsToLeaderboardData.
func (mr *MockMapperIMockRecorder) MapPSQLRowsToLeaderboardData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapPSQLRowsToLeaderboardData", reflect.TypeOf((*MockMapperI)(nil).MapPSQLRowsToLeaderboardData), arg0)
}
