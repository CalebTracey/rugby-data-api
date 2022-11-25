package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/request"
	"github.com/calebtracey/rugby-models/response"
	"strings"
)

const PSQLDatabaseSource = "rugby_db"

//go:generate mockgen -destination=../mocks/mockApiFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	GetCompetitionData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) GetCompetitionData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	//TODO add validation
	if strings.EqualFold(req.Source, PSQLDatabaseSource) {
		resp = s.CompService.LeaderboardData(ctx, req.CompName)
	}
	return resp
}
