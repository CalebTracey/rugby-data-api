package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	lbReq "github.com/calebtracey/rugby-models/pkg/dtos/request/leaderboard"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	lbRes "github.com/calebtracey/rugby-models/pkg/dtos/response/leaderboard"
)

//go:generate mockgen -destination=../mocks/mockApiFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	GetLeaderboardData(ctx context.Context, req lbReq.Request) (resp lbRes.Response)
	GetAllLeaderboardData(ctx context.Context) (resp lbRes.Response)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) GetLeaderboardData(ctx context.Context, req lbReq.Request) (resp lbRes.Response) {
	validationErrs := req.Validate()
	if validationErrs != nil {
		return lbRes.Response{
			Message: response.Message{ErrorLog: response.ErrorLogs{*validationErrs}},
		}
	}
	resp = s.CompService.LeaderboardData(ctx, req)

	return resp
}

func (s APIFacade) GetAllLeaderboardData(ctx context.Context) (resp lbRes.Response) {
	//TODO add validation
	resp = s.CompService.AllLeaderboardData(ctx)

	return resp
}
