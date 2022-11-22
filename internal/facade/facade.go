package facade

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/request"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/facade/sixnations"
	"strings"
)

type APIFacadeI interface {
	SixNationsResults(ctx context.Context, req request.CompetitionRequest) (resp response.CompetitionResponse)
}

type APIFacade struct {
	SNDAO sixnations.FacadeI
}

func (s APIFacade) SixNationsResults(ctx context.Context, req request.CompetitionRequest) (resp response.CompetitionResponse) {
	//TODO add validation
	if strings.EqualFold(req.Source, "DB") {
		resp = s.SNDAO.SixNationsTeams(ctx)
	}
	return resp
}
