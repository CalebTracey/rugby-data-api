package comp

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/calebtracey/rugby-models/pkg/models"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFacade_LeaderboardData(t *testing.T) {
	_, mock, _ := sqlmock.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCompDao := compmocks.NewMockDAOI(ctrl)
	mockCompMapper := compmocks.NewMockMapperI(ctrl)
	type fields struct {
		CompDAO    comp.DAOI
		CompMapper comp.MapperI
	}
	type args struct {
		ctx      context.Context
		compName string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantResp        response.LeaderboardResponse
		mockDaoResp     models.PSQLLeaderboardDataList
		wantCompId      string
		expectCompError bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				CompDAO:    mockCompDao,
				CompMapper: mockCompMapper,
			},
			args: args{
				ctx:      context.Background(),
				compName: "six nations",
			},
			wantCompId:      "180659",
			expectCompError: false,
			mockDaoResp: models.PSQLLeaderboardDataList{
				{
					CompId:   180659,
					CompName: "six nations",
					TeamId:   1,
					TeamName: "Team 1",
				},
				{
					CompId:   180659,
					CompName: "six nations",
					TeamId:   2,
					TeamName: "Team 2",
				},
			},
			wantResp: response.LeaderboardResponse{
				Id:   "180659",
				Name: "six nations",
				Teams: dtos.TeamLeaderboardDataList{
					{
						Id:   "1",
						Name: "Team 1",
					},
					{
						Id:   "2",
						Name: "Team 2",
					},
				},
				Message: response.Message{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Facade{
				CompDAO:    tt.fields.CompDAO,
				CompMapper: tt.fields.CompMapper,
			}
			query := fmt.Sprintf(comp.PSQLLeaderboardById, tt.wantCompId)
			mockCompMapper.EXPECT().CreatePSQLCompetitionQuery(tt.wantCompId).Return(query)
			mockCompDao.EXPECT().GetLeaderboardData(tt.args.ctx, query).
				DoAndReturn(func(ctx context.Context, query string) (models.PSQLLeaderboardDataList, *response.ErrorLog) {
					if tt.expectCompError {
						mock.ExpectQuery(query).WillReturnError(errors.New("db error"))
						return models.PSQLLeaderboardDataList{}, &response.ErrorLog{
							Query:      query,
							RootCause:  "db error",
							StatusCode: "500",
						}
					}
					return tt.mockDaoResp, nil
				})
			if !tt.expectCompError {
				mockCompMapper.EXPECT().MapPSQLLeaderboardDataToResponse(gomock.Any(), gomock.Any(), tt.mockDaoResp).
					Return(tt.wantResp)
			}
			if gotResp := s.LeaderboardData(tt.args.ctx, tt.args.compName); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
