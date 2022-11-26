package comp

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/calebtracey/rugby-models/pkg/models"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockDao.go -package=compmocks . DAOI
type DAOI interface {
	GetLeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog)
	GetAllLeaderboardData(ctx context.Context) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog)
}

type DAO struct {
	DbDAO    psql.DAOI
	DbMapper psql.MapperI
}

func (s DAO) GetLeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog) {
	rows, err := s.DbDAO.FindAll(ctx, query)
	if err != nil {
		return resp, err
	}
	resp = s.DbMapper.MapPSQLRowsToLeaderboardData(rows)

	return resp, nil
}

func (s DAO) GetAllLeaderboardData(ctx context.Context) (resp models.PSQLLeaderboardDataList, err *response.ErrorLog) {
	rows, err := s.DbDAO.FindAll(ctx, psql.AllLeaderboardsQuery)
	if err != nil {
		return resp, err
	}
	resp = s.DbMapper.MapPSQLRowsToLeaderboardData(rows)

	return resp, nil
}
