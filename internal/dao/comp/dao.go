package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockDao.go -package=compmocks . DAOI
type DAOI interface {
	GetLeaderboardData(ctx context.Context, query string) (psqlTeamsResponse response.LeaderboardResponse, err *response.ErrorLog)
}

type DAO struct {
	PSQLDAO    psql.DAOI
	PSQLMapper MapperI
}

func (s DAO) GetLeaderboardData(ctx context.Context, query string) (compResponse response.LeaderboardResponse, err *response.ErrorLog) {
	rows, err := s.PSQLDAO.FindAll(ctx, query)
	if err != nil {
		return compResponse, err
	}
	compData := s.PSQLMapper.MapPSQLRowsToCompetitionData(rows)
	compResponse = s.PSQLMapper.MapCompetitionDataResponse(compData)

	return compResponse, err
}
