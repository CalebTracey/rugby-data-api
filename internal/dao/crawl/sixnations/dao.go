package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	"github.com/gocolly/colly"
)

//go:generate mockgen -destination=mockDao.go -package=sixnations . DAOI
type DAOI interface {
	GetSixNationsData(ctx context.Context, url, date string) (resp response.TeamDataResponse, log *response.ErrorLog)
}

type DAO struct {
	Collector *colly.Collector
}

func (s DAO) GetSixNationsData(ctx context.Context, url, date string) (resp response.TeamDataResponse, log *response.ErrorLog) {
	//s.Collector.OnHTML()
	return resp, nil
}

func mapError(err error, query string) (errLog *response.ErrorLog) {
	errLog = &response.ErrorLog{
		Query: query,
	}
	if err != nil {
		errLog.RootCause = err.Error()
	}
	errLog.StatusCode = "500"
	return errLog
}
