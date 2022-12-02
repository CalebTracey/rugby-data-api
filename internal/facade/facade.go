package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
)

//go:generate mockgen -destination=../mocks/mockApiFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	GetLeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response)
	GetAllLeaderboardData(ctx context.Context) (resp leaderboard.Response)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) GetLeaderboardData(ctx context.Context, req leaderboard.Request) (resp leaderboard.Response) {
	validationErrs := req.Validate()
	if validationErrs != nil {
		return leaderboard.Response{
			Message: response.Message{ErrorLog: response.ErrorLogs{*validationErrs}},
		}
	}
	resp = s.CompService.LeaderboardData(ctx, req)

	return resp
}

func (s APIFacade) GetAllLeaderboardData(ctx context.Context) (resp leaderboard.Response) {
	resp = s.CompService.AllLeaderboardData(ctx)

	return resp
}
