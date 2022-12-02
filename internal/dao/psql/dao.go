package psql

import (
	"context"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/dbmocks/mockDao.go -package=dbmocks . DAOI
type DAOI interface {
	InsertOne(ctx context.Context, exec string) (res sql.Result, err error)
	FindAll(ctx context.Context, query string) (rows *sql.Rows, err error)
}

type DAO struct {
	DB *sql.DB
}

func (s DAO) InsertOne(ctx context.Context, exec string) (resp sql.Result, err error) {
	resp, sqlErr := s.DB.ExecContext(ctx, exec)
	if sqlErr != nil {
		log.Error(sqlErr)
		return resp, fmt.Errorf("error during leaderboard insert one: %w", sqlErr)
	}
	return resp, nil
}

func (s DAO) FindAll(ctx context.Context, query string) (rows *sql.Rows, err error) {
	rows, sqlErr := s.DB.QueryContext(ctx, query)
	if sqlErr != nil {
		log.Error(sqlErr)
		return rows, fmt.Errorf("error during leaderboard find all: %w", sqlErr)
	}
	return rows, nil
}
