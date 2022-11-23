package comp

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-data-api/external/models"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/psql"
	"github.com/golang/mock/gomock"
	"reflect"
	"regexp"
	"testing"
)

func TestSNDAO_GetTeams(t *testing.T) {
	_, mock, _ := sqlmock.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := psql.NewMockDAOI(ctrl)
	mockMapper := NewMockMapperI(ctrl)
	cols := []string{"comp_id", "comp_name", "team_id", "team_name"}

	type fields struct {
		PSQLDAO    psql.DAOI
		PSQLMapper MapperI
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockCols         []string
		mockCompDataList models.PSQLCompetitionDataList
		wantCompResponse response.CompetitionResponse
		wantErr          *response.ErrorLog
	}{
		{
			name: "Happy Path",
			fields: fields{
				PSQLDAO:    mockDao,
				PSQLMapper: mockMapper,
			},
			args: args{
				ctx:   context.Background(),
				query: fmt.Sprintf(PSQLCompetitionByID, "123"),
			},
			mockCols: cols,
			mockCompDataList: models.PSQLCompetitionDataList{
				{
					CompName: "Test",
					CompID:   123,
					TeamID:   321,
					TeamName: "Test",
				},
			},
			wantCompResponse: response.CompetitionResponse{
				ID:   "123",
				Name: "Test",
				Teams: models.TeamDataList{
					{
						Name: "Test",
						ID:   "321",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DAO{
				PSQLDAO:    tt.fields.PSQLDAO,
				PSQLMapper: tt.fields.PSQLMapper,
			}
			rows := sqlmock.NewRows(tt.mockCols).
				AddRow(1, "comp 1", 10, "team 1").
				AddRow(1, "comp 1", 5, "team 2")
			mock.ExpectBegin()
			mockDao.EXPECT().FindAll(tt.args.ctx, tt.args.query).
				DoAndReturn(func(ctx context.Context, query string) (*sqlmock.Rows, *response.ErrorLog) {
					mock.ExpectQuery(regexp.QuoteMeta(tt.args.query)).WillReturnRows(rows)
					return rows, nil
				})
			mockMapper.EXPECT().MapPSQLRowsToCompetitionData(gomock.Any()).Return(tt.mockCompDataList)
			mockMapper.EXPECT().MapCompetitionDataResponse(gomock.Any()).Return(tt.wantCompResponse)
			gotCompResponse, gotErr := s.GetTeams(tt.args.ctx, tt.args.query)
			if !reflect.DeepEqual(gotCompResponse, tt.wantCompResponse) {
				t.Errorf("GetTeams() gotCompResponse = %v, want %v", gotCompResponse, tt.wantCompResponse)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetTeams() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
