package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	log "github.com/sirupsen/logrus"
	"strings"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockFacade.go -package=compmocks . FacadeI
type FacadeI interface {
	LeaderboardData(ctx context.Context, compName string) (resp response.LeaderboardResponse)
	AllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse)
}

type Facade struct {
	CompDAO  comp.DAOI
	DbMapper psql.MapperI
}

func (s Facade) LeaderboardData(ctx context.Context, compName string) (resp response.LeaderboardResponse) {
	compId := getCompId(compName)
	teamsQuery := s.DbMapper.CreatePSQLLeaderboardByIdQuery(compId)

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
	resp = s.DbMapper.MapPSQLLeaderboardDataToResponse(compId, compName, leaderboardData)
	return resp
}

func (s Facade) AllLeaderboardData(ctx context.Context) (resp response.AllLeaderboardsResponse) {
	leaderboardData, err := s.CompDAO.GetAllLeaderboardData(ctx)
	if err != nil {
		log.Error(err)
		return response.AllLeaderboardsResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	resp = s.DbMapper.MapPSQLAllLeaderboardDataToResponse(leaderboardData)
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
