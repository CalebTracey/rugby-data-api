package psql

import (
	"context"
	"database/sql"
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
		//err = mapError(sqlErr, exec)
		return resp, err
	}
	return resp, nil
}

func (s DAO) FindAll(ctx context.Context, query string) (rows *sql.Rows, err error) {
	rows, sqlErr := s.DB.QueryContext(ctx, query)
	if sqlErr != nil {
		log.Error(sqlErr)
		//err = mapError(sqlErr, query)
		return rows, err
	}
	return rows, nil
}

//
//func mapError(err error, query string) (errLog *response.ErrorLog) {
//	errLog = &response.ErrorLog{
//		Query: query,
//	}
//	if err == sql.ErrNoRows {
//		errLog.RootCause = "Not found in database"
//		errLog.StatusCode = "404"
//		return errLog
//	}
//
//	if err != nil {
//		errLog.RootCause = err.Error()
//	}
//	errLog.StatusCode = "500"
//	return errLog
//}
