package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/rugby-data-api/internal/dao/comp"
	"github.com/calebtracey/rugby-data-api/internal/dao/psql"
	"github.com/calebtracey/rugby-data-api/internal/facade"
	c "github.com/calebtracey/rugby-data-api/internal/facade/comp"
	log "github.com/sirupsen/logrus"
)

func initializeDAO(config config.Config) (facade.APIFacadeI, error) {
	psqlDbConfig, err := config.GetDatabaseConfig("PSQL")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return facade.APIFacade{
		CompService: c.Facade{
			CompDAO: comp.DAO{
				DbDAO: psql.DAO{
					DB: psqlDbConfig.DB,
				},
				DbMapper: psql.Mapper{},
			},
			DbMapper: psql.Mapper{},
		},
	}, nil
}
