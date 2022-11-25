package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	log "github.com/sirupsen/logrus"
)

const LeagueIDSixNations = "180659"

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	LeaderboardData(ctx context.Context) (resp response.LeaderboardResponse)
}

type Facade struct {
	CompDAO    comp.DAOI
	CompMapper comp.MapperI
}

func (s Facade) LeaderboardData(ctx context.Context) (resp response.LeaderboardResponse) {
	teamsQuery := s.CompMapper.CreatePSQLCompetitionQuery(LeagueIDSixNations)

	resp, err := s.CompDAO.GetLeaderboardData(ctx, teamsQuery)
	if err != nil {
		log.Error(err)
		return response.LeaderboardResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	return resp
}
