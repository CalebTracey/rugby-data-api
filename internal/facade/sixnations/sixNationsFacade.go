package sixnations

import (
	"context"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	sncrawl "github.com/calebtracey/rugby-data-api/internal/dao/crawl/sixnations"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/sixnations"
	log "github.com/sirupsen/logrus"
)

const LeagueIDSixNations = "180659"

type FacadeI interface {
	SixNationsTeams(ctx context.Context) (resp response.CompetitionResponse)
}

type Facade struct {
	SNCrawler sncrawl.DAOI
	SNPSQL    sixnations.SNDAOI
	SNMapper  sixnations.MapperI
}

func (s Facade) SixNationsTeams(ctx context.Context) (resp response.CompetitionResponse) {
	teamsQuery := s.SNMapper.CreatePSQLTeamsQuery(LeagueIDSixNations)

	data, err := s.SNPSQL.GetTeams(ctx, teamsQuery)
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
	resp.Teams = data.Teams
	return resp
}
