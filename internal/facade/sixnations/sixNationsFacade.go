package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/request"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	sncrawl "github.com/calebtracey/rugby-data-api/internal/dao/crawl/sixnations"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/comp"
	log "github.com/sirupsen/logrus"
)

const LeagueIDSixNations = "180659"

//go:generate mockgen -destination=mockFacade.go -package=comp . FacadeI
type FacadeI interface {
	SixNationsTeams(ctx context.Context) (resp response.CompetitionResponse)
	SixNationsCrawl(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse)
}

type Facade struct {
	SNCrawler sncrawl.DAOI
	SNPSQL    comp.DAOI
	SNMapper  comp.MapperI
}

func (s Facade) SixNationsTeams(ctx context.Context) (resp response.CompetitionResponse) {
	teamsQuery := s.SNMapper.CreatePSQLCompetitionQuery(LeagueIDSixNations)

	resp, err := s.SNPSQL.GetTeams(ctx, teamsQuery)
	if err != nil {
		log.Error(err)
		return response.CompetitionResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	return resp
}

func (s Facade) SixNationsCrawl(ctx context.Context, req request.CompetitionCrawlRequest) (resp response.CompetitionCrawlResponse) {
	//TODO create scrape url
	url := ""
	_, err := s.SNCrawler.GetSixNationsData(ctx, url, req.Date)
	if err != nil {
		log.Error(err)
		return response.CompetitionCrawlResponse{
			Message: response.Message{
				ErrorLog: response.ErrorLogs{
					*err,
				},
			},
		}
	}
	//TODO add response mapping
	return resp
}
