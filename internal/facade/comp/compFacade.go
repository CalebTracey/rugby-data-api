package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"strings"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	LeaderboardData(ctx context.Context, compName string) (resp response.LeaderboardResponse)
}

type Facade struct {
	CompDAO    comp.DAOI
	CompMapper comp.MapperI
}

func (s Facade) LeaderboardData(ctx context.Context, compName string) (resp response.LeaderboardResponse) {
	compId := getCompId(compName)
	teamsQuery := s.CompMapper.CreatePSQLCompetitionQuery(compId)

	leaderboardData, err := s.CompDAO.GetLeaderboardData(ctx, teamsQuery)
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
	resp = s.CompMapper.MapPSQLLeaderboardDataToResponse(compId, compName, leaderboardData)
	return resp
}

func getCompId(compName string) string {
	switch strings.ToLower(compName) {
	case SixNations:
		return SixNationsId
	default:
		return ""
	}
}

const (
	SixNations   = "six nations"
	SixNationsId = "180659"
)
