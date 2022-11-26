package psql

import (
	"fmt"
	"testing"
)

func TestMapper_CreatePSQLCompetitionQuery(t *testing.T) {
	tests := []struct {
		name   string
		teamId string
		want   string
	}{
		{
			name:   "Happy Path",
			teamId: "123",
			want:   fmt.Sprintf(LeaderboardByIdQuery, 123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if got := m.CreatePSQLLeaderboardByIdQuery(tt.teamId); got != tt.want {
				t.Errorf("CreatePSQLLeaderboardByIdQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestMapper_MapPSQLRowsToCompetitionData(t *testing.T) {
//	_, mock, _ := sqlmock.New()
//	cols := []string{"comp_id", "comp_name", "team_id", "team_name"}
//	tests := []struct {
//		name               string
//		rows *sql.Rows
//		wantSixNationsData models.PSQLCompetitionDataList
//	}{
//		{
//			name: "Happy Path",
//
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := DbMapper{}
//			if gotSixNationsData := m.MapPSQLRowsToLeaderboardData(tt.rows); !reflect.DeepEqual(gotSixNationsData, tt.wantSixNationsData) {
//				t.Errorf("MapPSQLRowsToLeaderboardData() = %v, want %v", gotSixNationsData, tt.wantSixNationsData)
//			}
//		})
//	}
//}
