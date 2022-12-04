package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	"github.com/calebtracey/rugby-data-api/internal/mocks/compmocks"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestAPIFacade_GetLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFacade := compmocks.NewMockFacadeI(ctrl)
	tests := []struct {
		name        string
		CompService comp.FacadeI
		ctx         context.Context
		req         leaderboard.Request
		wantResp    leaderboard.Response
	}{
		{
			name:        "Happy Path",
			CompService: mockFacade,
			ctx:         context.Background(),
			req: leaderboard.Request{
				Competitions: dtos.CompetitionList{
					{
						Name: "six nations",
					},
				},
			},
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompName: "Six Nations",
						CompId:   "180659",
						Teams:    dtos.TeamLeaderboardDataList{},
					},
				},
				Message: response.Message{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := APIFacade{
				CompService: tt.CompService,
			}

			mockFacade.EXPECT().LeaderboardData(tt.ctx, tt.req).Return(tt.wantResp)

			if gotResp := s.GetLeaderboardData(tt.ctx, tt.req); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetLeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestAPIFacade_GetAllLeaderboardData(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFacade := compmocks.NewMockFacadeI(ctrl)

	tests := []struct {
		name        string
		CompService comp.FacadeI
		ctx         context.Context

		wantResp leaderboard.Response
	}{
		{
			name:        "Happy Path",
			CompService: mockFacade,
			ctx:         context.Background(),
			wantResp: leaderboard.Response{
				LeaderboardData: dtos.CompetitionLeaderboardDataList{
					{
						CompName: "Six Nations",
						CompId:   "180659",
						Teams:    dtos.TeamLeaderboardDataList{},
					},
				},
				Message: response.Message{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := APIFacade{
				CompService: tt.CompService,
			}

			mockFacade.EXPECT().AllLeaderboardData(tt.ctx).Return(tt.wantResp)

			if gotResp := s.GetAllLeaderboardData(tt.ctx); !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetAllLeaderboardData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
