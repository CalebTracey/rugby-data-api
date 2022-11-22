package crawl

//
//import (
//	"context"
//	"github.com/calebtracey/rugby-data-api/external/models/response"
//	"github.com/gocolly/colly"
//)
//
//type DAOI interface {
//	GetTeamData(ctx context.Context, team string) (resp response.CrawlTeamResponse, log *response.ErrorLog)
//}
//
//type DAO struct {
//	Collector *colly.Collector
//}
//
//func (s DAO) GetTeamData(ctx context.Context, team string) (resp response.CrawlTeamResponse, log *response.ErrorLog) {
//
//	return resp, nil
//}
//
//func mapError(err error, query string) (errLog *response.ErrorLog) {
//	errLog = &response.ErrorLog{
//		Query: query,
//	}
//	if err != nil {
//		errLog.RootCause = err.Error()
//	}
//	errLog.StatusCode = "500"
//	return errLog
//}
