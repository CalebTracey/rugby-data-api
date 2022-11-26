package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"strings"
)

const PSQLDatabaseSource = "rugby_db"

//go:generate mockgen -destination=../mocks/mockApiFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	GetLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
	GetAllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) GetLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	//TODO add validation
	if strings.EqualFold(req.Source, PSQLDatabaseSource) {
		resp = s.CompService.LeaderboardData(ctx, req.CompName)
	}
	return resp
}

func (s APIFacade) GetAllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse) {
	//TODO add validation
	resp = s.CompService.AllLeaderboardData(ctx)

	return resp
}
