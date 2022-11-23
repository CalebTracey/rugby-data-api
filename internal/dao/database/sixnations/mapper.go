package sixnations

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-data-api/external/models"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mockMapper.go -package=sixnations . MapperI
type MapperI interface {
	CreatePSQLTeamsQuery(teamId string) string
	MapPSQLRowsToTeamData(rows *sql.Rows) (sixNationsTeams []models.TeamData)
}

type Mapper struct{}

func (m Mapper) CreatePSQLTeamsQuery(teamId string) string {
	return fmt.Sprintf(PSQLSelectByID, teamId)
}

func (m Mapper) MapPSQLRowsToTeamData(rows *sql.Rows) (sixNationsTeams []models.TeamData) {
	var team models.TeamData
	for rows.Next() {
		err := rows.Scan(
			&team.Name,
			pq.Array(&team.Players),
		)
		if err != nil {
			log.Panicln(err)
		}
		sixNationsTeams = append(sixNationsTeams, team)
	}

	return sixNationsTeams
}

const PSQLSelectByID = `select id, name, players from public.sixnations.teams where id = '%s'`
