package comp

import (
	"context"
	"fmt"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-data-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-data-api/internal/mocks/dbmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/calebtracey/rugby-models/pkg/models"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestFacade_LeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCompDao := compmocks.NewMockDAOI(ctrl)
	mockCompMapper := dbmocks.NewMockMapperI(ctrl)
	mockQuery := fmt.Sprintf(psql.LeaderboardByIdQuery, "180659")
	type fields struct {
		CompDAO  comp.DAOI
		DbMapper psql.MapperI
	}
	type args struct {
		ctx      context.Context
		compName string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		query         string
		wantResp      response.LeaderboardResponse
		mockDaoResp   models.PSQLLeaderboardDataList
		wantCompId    string
		expectDbError bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				CompDAO:  mockCompDao,
				DbMapper: mockCompMapper,
			},
			args: args{
				ctx:      context.Background(),
				compName: "six nations",
			},
			query:         mockQuery,
			wantCompId:    "180659",
			expectDbError: false,
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
				LeaderboardData: dtos.CompetitionLeaderboardData{
					CompId:   "180659",
					CompName: "Six Nations",
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
				},
				Message: response.Message{},
			},
		},
		{
			name: "Sad Path",
			fields: fields{
				CompDAO:  mockCompDao,
				DbMapper: mockCompMapper,
			},
			args: args{
				ctx:      context.Background(),
				compName: "six nations",
			},
			wantCompId:    "180659",
			expectDbError: true,
			query:         mockQuery,
			wantResp: response.LeaderboardResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query:      mockQuery,
							RootCause:  "db error",
							StatusCode: "500",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Facade{
				CompDAO:  tt.fields.CompDAO,
				DbMapper: tt.fields.DbMapper,
			}
			mockCompMapper.EXPECT().CreatePSQLLeaderboardByIdQuery(tt.wantCompId).Return(tt.query)
			mockCompDao.EXPECT().GetLeaderboardData(tt.args.ctx, tt.query).
				DoAndReturn(func(ctx context.Context, query string) (models.PSQLLeaderboardDataList, *response.ErrorLog) {
					if tt.expectDbError {
						return models.PSQLLeaderboardDataList{}, &response.ErrorLog{
							Query:      tt.query,
							RootCause:  "db error",
							StatusCode: "500",
						}
					}
					return tt.mockDaoResp, nil
				})
			if !tt.expectDbError {
				mockCompMapper.EXPECT().MapPSQLLeaderboardDataToResponse(gomock.Any(), gomock.Any(), tt.mockDaoResp).
					Return(tt.wantResp)
			}
			if gotResp := s.LeaderboardData(tt.args.ctx, tt.args.compName); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestFacade_AllLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCompDao := compmocks.NewMockDAOI(ctrl)
	mockCompMapper := dbmocks.NewMockMapperI(ctrl)
	type fields struct {
		CompDAO  comp.DAOI
		DbMapper psql.MapperI
	}
	tests := []struct {
		name          string
		fields        fields
		ctx           context.Context
		query         string
		wantResp      response.AllLeaderboardsResponse
		mockDaoResp   models.PSQLLeaderboardDataList
		mockDaoErr    *response.ErrorLog
		expectDbError bool
	}{
		{
			name: "Happy Path",
			fields: fields{
				CompDAO:  mockCompDao,
				DbMapper: mockCompMapper,
			},
			ctx:           context.Background(),
			expectDbError: false,
			mockDaoResp: models.PSQLLeaderboardDataList{
				{
					CompId:   180659,
					CompName: "Six Nations",
					TeamId:   1,
					TeamName: "Team 1",
				},
				{
					CompId:   180659,
					CompName: "Six Nations",
					TeamId:   2,
					TeamName: "Team 2",
				},
				{
					CompId:   267979,
					CompName: "Premiership",
					TeamId:   3,
					TeamName: "Team 3",
				},
				{
					CompId:   267979,
					CompName: "Premiership",
					TeamId:   4,
					TeamName: "Team 4",
				},
			},
			mockDaoErr: nil,
			wantResp: response.AllLeaderboardsResponse{
				LeaderboardDataList: dtos.CompetitionLeaderboardDataList{
					{
						CompId:   SixNationsId,
						CompName: SixNations,
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
					},
					{
						CompId:   PremiershipId,
						CompName: Premiership,
						Teams: dtos.TeamLeaderboardDataList{
							{
								Id:   "3",
								Name: "Team 3",
							},
							{
								Id:   "4",
								Name: "Team 4",
							},
						},
					},
				},
			},
		},
		{
			name: "Sad Path",
			fields: fields{
				CompDAO:  mockCompDao,
				DbMapper: mockCompMapper,
			},
			ctx:           context.Background(),
			expectDbError: true,
			query:         psql.AllLeaderboardsQuery,
			wantResp: response.AllLeaderboardsResponse{
				Message: response.Message{
					ErrorLog: response.ErrorLogs{
						{
							Query:      psql.AllLeaderboardsQuery,
							RootCause:  "db error",
							StatusCode: "500",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Facade{
				CompDAO:  tt.fields.CompDAO,
				DbMapper: tt.fields.DbMapper,
			}
			mockCompDao.EXPECT().GetAllLeaderboardData(tt.ctx).
				DoAndReturn(func(ctx context.Context) (models.PSQLLeaderboardDataList, *response.ErrorLog) {
					if tt.expectDbError {
						return models.PSQLLeaderboardDataList{}, &response.ErrorLog{
							Query:      tt.query,
							RootCause:  "db error",
							StatusCode: "500",
						}
					}
					return tt.mockDaoResp, nil
				})
			if !tt.expectDbError {
				mockCompMapper.EXPECT().MapPSQLAllLeaderboardDataToResponse(tt.mockDaoResp).Return(tt.wantResp)
			}
			if gotResp := s.AllLeaderboardData(tt.ctx); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AllLeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
