package sixnations

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-data-api/external/models"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mockMapper.go -package=sixnations . MapperI
type MapperI interface {
	CreatePSQLCompetitionQuery(teamId string) string
	MapPSQLRowsToCompetitionData(rows *sql.Rows) (sixNationsData models.PSQLCompetitionDataList)
}

type Mapper struct{}

func (m Mapper) CreatePSQLCompetitionQuery(teamId string) string {
	return fmt.Sprintf(PSQLCompetitionByID, teamId)
}

func (m Mapper) MapPSQLRowsToCompetitionData(rows *sql.Rows) (sixNationsData models.PSQLCompetitionDataList) {
	var data models.PSQLCompetitionData
	for rows.Next() {
		err := rows.Scan(
			&data.CompID,
			&data.CompName,
			&data.TeamID,
			&data.TeamName,
		)
		if err != nil {
			log.Panicln(err)
		}
		sixNationsData = append(sixNationsData, data)
	}

	return sixNationsData
}

const (
	PSQLCompetitionByID = `with comp_teams as (
		select
			c.comp_id,
			c.name as comp_name,
			t.name as team_name,
			t.team_id
		from
			competitions c,
			unnest(c.teams) t_id
			join teams t on (t.team_id = t_id)
		where
			c.comp_id = '%s'
	) select comp_id, comp_name, team_id, team_name from comp_teams;`
)
