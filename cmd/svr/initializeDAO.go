package main

import (
	config "github.com/calebtracey/config-yaml"
	psql2 "github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	"github.com/calebtracey/rugby-data-api/internal/facade/psql"
	log "github.com/sirupsen/logrus"
)

func initializeDAO(config config.Config) (facade.APIFacadeI, error) {
	psqlDbConfig, err := config.GetDatabaseConfig("PSQL")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return facade.APIFacade{
		PSQLDao: psql.Facade{
			PSQL: psql2.DAO{
				DB: psqlDbConfig.DB,
			},
			PSQLMapper: psql2.Mapper{},
		},
	}, nil
}
