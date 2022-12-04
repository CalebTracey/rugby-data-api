package psql

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
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
			if got := m.LeaderboardByIdQuery(tt.teamId); got != tt.want {
				t.Errorf("LeaderboardByIdQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapper_MapPSQLRowsToLeaderboardData(t *testing.T) {
	db, mock, _ := sqlmock.New()
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
		wantErr             error
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
			wantErr: nil,
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
			gotLeaderboardData, gotErrorLog := m.RowsToLeaderboardData(rows)
			if !reflect.DeepEqual(gotLeaderboardData, tt.wantLeaderboardData) {
				t.Errorf("RowsToLeaderboardData() gotLeaderboardData = %v, want %v", gotLeaderboardData, tt.wantLeaderboardData)
			}
			if !reflect.DeepEqual(gotErrorLog, tt.wantErr) {
				t.Errorf("RowsToLeaderboardData() gotErrorLog = %v, want %v", gotErrorLog, tt.wantErr)
			}
		})
	}
}

func Test_mapPSQLTeamData(t *testing.T) {

	tests := []struct {
		name string
		data models.PSQLLeaderboardData
		want dtos.TeamLeaderboardData
	}{
		{
			name: "Happy Path",
			data: models.PSQLLeaderboardData{
				CompId:            123,
				CompName:          "Comp",
				TeamId:            321,
				TeamName:          "Team",
				TeamAbbr:          "ABBR",
				GamesPlayed:       "3",
				WinCount:          "4",
				DrawCount:         "5",
				LossCount:         "6",
				Bye:               "7",
				PointsFor:         "8",
				PointsAgainst:     "9",
				TriesFor:          "10",
				TriesAgainst:      "11",
				BonusPointsTry:    "12",
				BonusPointsLosing: "13",
				BonusPoints:       "14",
				PointsDiff:        "15",
				Points:            "16",
			},
			want: dtos.TeamLeaderboardData{
				Id:   "321",
				Name: "Team",
				Abbr: "ABBR",
				CompetitionStats: dtos.TeamCompetitionStats{
					GamesPlayed:       "3",
					WinCount:          "4",
					DrawCount:         "5",
					LossCount:         "6",
					Bye:               "7",
					PointsFor:         "8",
					PointsAgainst:     "9",
					TriesFor:          "10",
					TriesAgainst:      "11",
					BonusPointsTry:    "12",
					BonusPointsLosing: "13",
					BonusPoints:       "14",
					PointsDiff:        "15",
					Points:            "16",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapPSQLTeamData(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapPSQLTeamData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapper_LeaderboardDataToResponse(t *testing.T) {
	type args struct {
		compId          string
		compName        string
		leaderboardData models.PSQLLeaderboardDataList
	}
	tests := []struct {
		name     string
		args     args
		wantResp dtos.CompetitionLeaderboardData
	}{
		{
			name: "Happy Path",
			args: args{
				compId:   "123",
				compName: "321",
				leaderboardData: models.PSQLLeaderboardDataList{
					{
						CompId:            123,
						CompName:          "Comp",
						TeamId:            321,
						TeamName:          "Team",
						TeamAbbr:          "ABBR",
						GamesPlayed:       "3",
						WinCount:          "4",
						DrawCount:         "5",
						LossCount:         "6",
						Bye:               "7",
						PointsFor:         "8",
						PointsAgainst:     "9",
						TriesFor:          "10",
						TriesAgainst:      "11",
						BonusPointsTry:    "12",
						BonusPointsLosing: "13",
						BonusPoints:       "14",
						PointsDiff:        "15",
						Points:            "16",
					},
				},
			},
			wantResp: dtos.CompetitionLeaderboardData{
				CompId:   "123",
				CompName: "321",
				Teams: dtos.TeamLeaderboardDataList{
					{
						Id:   "321",
						Name: "Team",
						Abbr: "ABBR",
						CompetitionStats: dtos.TeamCompetitionStats{
							GamesPlayed:       "3",
							WinCount:          "4",
							DrawCount:         "5",
							LossCount:         "6",
							Bye:               "7",
							PointsFor:         "8",
							PointsAgainst:     "9",
							TriesFor:          "10",
							TriesAgainst:      "11",
							BonusPointsTry:    "12",
							BonusPointsLosing: "13",
							BonusPoints:       "14",
							PointsDiff:        "15",
							Points:            "16",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if gotResp := m.LeaderboardDataToResponse(tt.args.compId, tt.args.compName, tt.args.leaderboardData); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LeaderboardDataToResponse() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestMapper_AllLeaderboardDataToResponse(t *testing.T) {

	tests := []struct {
		name                string
		leaderboardDataList models.PSQLLeaderboardDataList
		wantResp            leaderboard.Response
	}{
		{
			name: "Happy Path",
			leaderboardDataList: models.PSQLLeaderboardDataList{
				{
					CompId:            123,
					CompName:          "Comp",
					TeamId:            321,
					TeamName:          "Team",
					TeamAbbr:          "ABBR",
					GamesPlayed:       "3",
					WinCount:          "4",
					DrawCount:         "5",
					LossCount:         "6",
					Bye:               "7",
					PointsFor:         "8",
					PointsAgainst:     "9",
					TriesFor:          "10",
					TriesAgainst:      "11",
					BonusPointsTry:    "12",
					BonusPointsLosing: "13",
					BonusPoints:       "14",
					PointsDiff:        "15",
					Points:            "16",
				},
			},
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompId:   "123",
						CompName: "Comp",
						Teams: dtos.TeamLeaderboardDataList{
							{
								Id:   "321",
								Name: "Team",
								Abbr: "ABBR",
								CompetitionStats: dtos.TeamCompetitionStats{
									GamesPlayed:       "3",
									WinCount:          "4",
									DrawCount:         "5",
									LossCount:         "6",
									Bye:               "7",
									PointsFor:         "8",
									PointsAgainst:     "9",
									TriesFor:          "10",
									TriesAgainst:      "11",
									BonusPointsTry:    "12",
									BonusPointsLosing: "13",
									BonusPoints:       "14",
									PointsDiff:        "15",
									Points:            "16",
								},
							},
						},
					},
				},
				Message: response.Message{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mapper{}
			if gotResp := m.AllLeaderboardDataToResponse(tt.leaderboardDataList); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AllLeaderboardDataToResponse() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
