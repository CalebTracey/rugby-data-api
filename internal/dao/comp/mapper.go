package comp

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-data-api/external/models"
	"github.com/calebtracey/rugby-data-api/external/models/response"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockMapper.go -package=compmocks . MapperI
type MapperI interface {
	CreatePSQLCompetitionQuery(teamId string) string
	MapPSQLRowsToCompetitionData(rows *sql.Rows) (sixNationsData models.PSQLCompetitionDataList)
	MapCompetitionDataResponse(sixNationsData models.PSQLCompetitionDataList) (resp response.LeaderboardResponse)
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

func (m Mapper) MapCompetitionDataResponse(sixNationsData models.PSQLCompetitionDataList) (resp response.LeaderboardResponse) {
	resp.ID = strconv.Itoa(sixNationsData[0].CompID)
	resp.Name = sixNationsData[0].CompName
	for _, data := range sixNationsData {
		resp.Teams = append(resp.Teams, models.TeamData{
			ID:   strconv.Itoa(data.TeamID),
			Name: data.TeamName,
		})
	}
	return resp
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
