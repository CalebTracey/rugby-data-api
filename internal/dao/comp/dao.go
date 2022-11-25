package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/models"
	"github.com/calebtracey/rugby-models/response"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockDao.go -package=compmocks . DAOI
type DAOI interface {
	GetLeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog)
}

type DAO struct {
	DbDAO  psql.DAOI
	Mapper MapperI
}

func (s DAO) GetLeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog) {
	rows, err := s.DbDAO.FindAll(ctx, query)
	if err != nil {
		return resp, err
	}
	resp = s.Mapper.MapPSQLRowsToLeaderboardData(rows)

	return resp, nil
}
