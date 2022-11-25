package comp

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-data-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-data-api/internal/mocks/dbmocks"
	"github.com/calebtracey/rugby-models/models"
	"github.com/calebtracey/rugby-models/response"
	"github.com/golang/mock/gomock"
	"reflect"
	"regexp"
	"testing"
)

func TestDAO_GetLeaderboardData(t *testing.T) {
	_, mock, _ := sqlmock.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := dbmocks.NewMockDAOI(ctrl)
	mockMapper := compmocks.NewMockMapperI(ctrl)
	cols := []string{"comp_id", "comp_name", "team_id", "team_name"}

	type fields struct {
		DbDAO  psql.DAOI
		Mapper MapperI
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockCols []string
		wantRes  models.PSQLLeaderboardDataList
		wantErr  *response.ErrorLog
	}{
		{
			name: "Happy Path",
			fields: fields{
				DbDAO:  mockDao,
				Mapper: mockMapper,
			},
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(PSQLCompetitionByID, "123"),
			},
			mockCols: cols,
			wantRes: models.PSQLLeaderboardDataList{
				{
					CompId:   "123",
					CompName: "Test Comp",
					TeamId:   "1",
					TeamName: "Team 1",
				},
				{
					CompId:   "123",
					CompName: "Test Comp",
					TeamId:   "2",
					TeamName: "Team 2",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				DbDAO:  tt.fields.DbDAO,
				Mapper: tt.fields.Mapper,
			}
			rows := sqlmock.NewRows(tt.mockCols).
				AddRow(123, "Test Comp", 1, "Team 1").
				AddRow(123, "Test Comp", 2, "Team 2")
			mock.ExpectBegin()
			mockDao.EXPECT().FindAll(tt.args.ctx, tt.args.query).
				DoAndReturn(func(ctx context.Context, query string) (*sqlmock.Rows, *response.ErrorLog) {
					mock.ExpectQuery(regexp.QuoteMeta(tt.args.query)).WillReturnRows(rows)
					return rows, nil
				})
			mockMapper.EXPECT().MapPSQLRowsToLeaderboardData(gomock.Any()).Return(tt.wantRes)
			gotCompResponse, gotErr := s.GetLeaderboardData(tt.args.ctx, tt.args.query)
			if !reflect.DeepEqual(gotCompResponse, tt.wantRes) {
				t.Errorf("GetLeaderboardData() gotCompResponse = %v, want %v", gotCompResponse, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetLeaderboardData() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
