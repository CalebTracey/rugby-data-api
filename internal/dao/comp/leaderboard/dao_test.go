package leaderboard

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-data-api/internal/mocks/dbmocks"
	"github.com/calebtracey/rugby-models/pkg/models"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestDAO_GetLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dbmocks.NewMockDAOI(ctrl)
	mockMapper := dbmocks.NewMockMapperI(ctrl)
	mockRows := sqlmock.NewRows([]string{"comp_id", "comp_name", "team_id", "team_name"}).
		AddRow(123, "Test Comp", 1, "Team 1").
		AddRow(123, "Test Comp", 2, "Team 2")
	type fields struct {
		DbDAO    psql.DAOI
		DbMapper psql.MapperI
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockRows      *sqlmock.Rows
		wantRes       models.PSQLLeaderboardDataList
		wantErr       error
		mockErr       error
		wantMapErr    error
		expectDbError bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				DbDAO:    mockDao,
				DbMapper: mockMapper,
			},
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(psql.CompetitionByIdQuery, "123"),
			},
			mockRows: mockRows,
			wantRes: models.PSQLLeaderboardDataList{
				{
					CompId:   123,
					CompName: "Test Comp",
					TeamId:   1,
					TeamName: "Team 1",
				},
				{
					CompId:   123,
					CompName: "Test Comp",
					TeamId:   2,
					TeamName: "Team 2",
				},
			},
			mockErr:       nil,
			wantErr:       nil,
			wantMapErr:    nil,
			expectDbError: false,
		},
		{
			name: "Sad Path - db error",
			fields: fields{
				DbDAO:    mockDao,
				DbMapper: mockMapper,
			},
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(psql.LeaderboardByIdQuery, "123"),
			},
			mockRows:      mockRows,
			wantRes:       models.PSQLLeaderboardDataList(nil),
			mockErr:       errors.New("db error"),
			wantErr:       errors.New("db error"),
			wantMapErr:    nil,
			expectDbError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DbDAO:    tt.fields.DbDAO,
				DbMapper: tt.fields.DbMapper,
			}
			mockDao.EXPECT().FindAll(tt.args.ctx, tt.args.query).
				DoAndReturn(func(ctx context.Context, query string) (*sqlmock.Rows, error) {
					if tt.expectDbError {
						return tt.mockRows, tt.wantErr
					}
					return tt.mockRows, nil
				})
			if !tt.expectDbError {
				mockMapper.EXPECT().RowsToLeaderboardData(gomock.Any()).Return(tt.wantRes, tt.wantMapErr)
			}
			gotCompResponse, gotErr := s.LeaderboardData(tt.args.ctx, tt.args.query)
			if !reflect.DeepEqual(gotCompResponse, tt.wantRes) {
				t.Errorf("LeaderboardData() gotCompResponse = %v, want %v", gotCompResponse, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("LeaderboardData() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestDAO_AllLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dbmocks.NewMockDAOI(ctrl)
	mockMapper := dbmocks.NewMockMapperI(ctrl)
	mockRows := sqlmock.NewRows([]string{"comp_id", "comp_name", "team_id", "team_name"}).
		AddRow(123, "Test Comp", 1, "Team 1").
		AddRow(123, "Test Comp", 2, "Team 2")
	type fields struct {
		DbDAO    psql.DAOI
		DbMapper psql.MapperI
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		mockRows      *sqlmock.Rows
		wantResp      models.PSQLLeaderboardDataList
		wantErr       error
		mockErr       error
		wantMapErr    error
		expectDbError bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				DbDAO:    mockDao,
				DbMapper: mockMapper,
			},
			args: args{
				ctx: context.Background(),
			},
			mockRows: mockRows,
			wantResp: models.PSQLLeaderboardDataList{
				{
					CompId:   123,
					CompName: "Test Comp",
					TeamId:   1,
					TeamName: "Team 1",
				},
				{
					CompId:   123,
					CompName: "Test Comp",
					TeamId:   2,
					TeamName: "Team 2",
				},
			},
			mockErr:       nil,
			wantErr:       nil,
			wantMapErr:    nil,
			expectDbError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DbDAO:    tt.fields.DbDAO,
				DbMapper: tt.fields.DbMapper,
			}
			mockDao.EXPECT().FindAll(tt.args.ctx, psql.AllLeaderboardsQuery).
				DoAndReturn(func(ctx context.Context, query string) (*sqlmock.Rows, error) {
					if tt.expectDbError {
						return tt.mockRows, tt.wantErr
					}
					return tt.mockRows, nil
				})
			if !tt.expectDbError {
				mockMapper.EXPECT().RowsToLeaderboardData(gomock.Any()).Return(tt.wantResp, tt.wantMapErr)
			}
			gotResp, err := s.AllLeaderboardData(tt.args.ctx)
			if err != tt.wantErr {
				t.Errorf("AllLeaderboardData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AllLeaderboardData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
