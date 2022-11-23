package main

import (
	config "github.com/calebtracey/config-yaml"
	sncrawl "github.com/calebtracey/rugby-data-api/internal/dao/crawl/sixnations"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/psql"
	"github.com/calebtracey/rugby-data-api/internal/dao/database/sixnations"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	sn "github.com/calebtracey/rugby-data-api/internal/facade/sixnations"
	log "github.com/sirupsen/logrus"
)

func initializeDAO(config config.Config) (facade.APIFacadeI, error) {
	psqlDbConfig, err := config.GetDatabaseConfig("PSQL")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	collyCrawlerConfig, err := config.GetCrawlConfig("COLLY")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return facade.APIFacade{
		SNDAO: sn.Facade{
			SNCrawler: sncrawl.DAO{
				Collector: collyCrawlerConfig.Collector,
			},
			SNPSQL: sixnations.SNDAO{
				PSQLDAO: psql.DAO{
					DB: psqlDbConfig.DB,
				},
				PSQLMapper: sixnations.Mapper{},
			},
			SNMapper: sixnations.Mapper{},
		},
	}, nil
}
