package leaderboard

import (
	"context"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-models/pkg/models"
)

//go:generate mockgen -destination=../../../mocks/compmocks/mockDao.go -package=compmocks . DAOI
type DAOI interface {
	LeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err error)
	AllLeaderboardData(ctx context.Context) (resp models.PSQLLeaderboardDataList, err error)
}

type DAO struct {
	DbDAO    psql.DAOI
	DbMapper psql.MapperI
}

func (s DAO) LeaderboardData(ctx context.Context, query string) (resp models.PSQLLeaderboardDataList, err error) {
	rows, err := s.DbDAO.FindAll(ctx, query)
	if err != nil {
		return resp, err
	}
	resp, err = s.DbMapper.RowsToLeaderboardData(rows)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s DAO) AllLeaderboardData(ctx context.Context) (resp models.PSQLLeaderboardDataList, err error) {
	rows, err := s.DbDAO.FindAll(ctx, psql.AllLeaderboardsQuery)
	if err != nil {
		return resp, err
	}
	resp, err = s.DbMapper.RowsToLeaderboardData(rows)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
