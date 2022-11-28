package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/cmd/validate"
	"github.com/calebtracey/rugby-data-api/internal/facade/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/request"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
)

const PSQLDatabaseSource = "rugby_db"

//go:generate mockgen -destination=../mocks/mockApiFacade.go -package=mocks . APIFacadeI
type APIFacadeI interface {
	GetLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse)
	GetAllLeaderboardData(ctx context.Context) (resp response.LeaderboardResponse)
}

type APIFacade struct {
	CompService comp.FacadeI
}

func (s APIFacade) GetLeaderboardData(ctx context.Context, req request.LeaderboardRequest) (resp response.LeaderboardResponse) {
	//TODO add validation
	validationErrs := validate.StructValidation(req)
	if validationErrs != nil {
		return response.LeaderboardResponse{
			Message: response.Message{ErrorLog: validationErrs},
		}
	}
	resp = s.CompService.LeaderboardData(ctx, req)

	return resp
}

func (s APIFacade) GetAllLeaderboardData(ctx context.Context) (resp response.LeaderboardResponse) {
	//TODO add validation
	resp = s.CompService.AllLeaderboardData(ctx)

	return resp
}
