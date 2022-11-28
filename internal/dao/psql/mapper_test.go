package psql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/calebtracey/rugby-models/pkg/models"
	log "github.com/sirupsen/logrus"
	"reflect"
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

var ()

func TestMapper_MapPSQLRowsToLeaderboardData(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error(err)
		}
	}(db)
	mockCols := []string{"comp_id",
		"comp_name",
		"team_id",
		"team_name",
		"team_abbr",
		"games_played",
		"win_count",
		"draw_count",
		"loss_count",
		"bye",
		"points_for",
		"points_against",
		"tries_for",
		"tries_against",
		"bonus_points_try",
		"bonus_points_losing",
		"bonus_points",
		"points_diff",
		"points"}
	mockRows := sqlmock.NewRows(mockCols).
		AddRow(123, "Test Comp", 1, "Team 1", "T1", "2", "2", "2", "2", "2", "2", "2", "6", "23", "51", "3", "4", "5", "15").
		AddRow(123, "Test Comp", 2, "Team 2", "T2", "2", "2", "2", "2", "2", "2", "2", "6", "23", "51", "3", "4", "5", "15")

	tests := []struct {
		name                string
		wantLeaderboardData models.PSQLLeaderboardDataList
		wantErrorLog        *response.ErrorLog
	}{
		{
			name: "Happy Path",
			wantLeaderboardData: models.PSQLLeaderboardDataList{
				{
					CompId:            123,
					CompName:          "Test Comp",
					TeamId:            1,
					TeamName:          "Team 1",
					TeamAbbr:          "T1",
					GamesPlayed:       "2",
					WinCount:          "2",
					DrawCount:         "2",
					LossCount:         "2",
					Bye:               "2",
					PointsFor:         "2",
					PointsAgainst:     "2",
					TriesFor:          "6",
					TriesAgainst:      "23",
					BonusPointsTry:    "51",
					BonusPointsLosing: "3",
					BonusPoints:       "4",
					PointsDiff:        "5",
					Points:            "15",
				},
				{
					CompId:            123,
					CompName:          "Test Comp",
					TeamId:            2,
					TeamName:          "Team 2",
					TeamAbbr:          "T2",
					GamesPlayed:       "2",
					WinCount:          "2",
					DrawCount:         "2",
					LossCount:         "2",
					Bye:               "2",
					PointsFor:         "2",
					PointsAgainst:     "2",
					TriesFor:          "6",
					TriesAgainst:      "23",
					BonusPointsTry:    "51",
					BonusPointsLosing: "3",
					BonusPoints:       "4",
					PointsDiff:        "5",
					Points:            "15",
				},
			},
			wantErrorLog: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			mock.ExpectQuery(``).WillReturnRows(mockRows)
			rows, err := db.Query(``)
			if err != nil {
				log.Error(err)
			}
			gotLeaderboardData, gotErrorLog := m.MapPSQLRowsToLeaderboardData(rows)
			if !reflect.DeepEqual(gotLeaderboardData, tt.wantLeaderboardData) {
				t.Errorf("MapPSQLRowsToLeaderboardData() gotLeaderboardData = %v, want %v", gotLeaderboardData, tt.wantLeaderboardData)
			}
			if !reflect.DeepEqual(gotErrorLog, tt.wantErrorLog) {
				t.Errorf("MapPSQLRowsToLeaderboardData() gotErrorLog = %v, want %v", gotErrorLog, tt.wantErrorLog)
			}
		})
	}
}
