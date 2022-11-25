package comp

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-models/models"
	"github.com/calebtracey/rugby-models/response"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//go:generate mockgen -destination=../../mocks/compmocks/mockMapper.go -package=compmocks . MapperI
type MapperI interface {
	CreatePSQLCompetitionQuery(teamId string) string
	MapPSQLRowsToLeaderboardData(rows *sql.Rows) (leaderboardData models.PSQLLeaderboardDataList)
	MapPSQLLeaderboardDataToResponse(compId, compName string, leaderboardData models.PSQLLeaderboardDataList) (resp response.LeaderboardResponse)
}

type Mapper struct{}

func (m Mapper) CreatePSQLCompetitionQuery(teamId string) string {
	teamIdInt, err := strconv.Atoi(teamId)
	if err != nil {
		log.Error(err)
		return ""
	}
	return fmt.Sprintf(PSQLLeaderboardById, teamIdInt)
}

func (m Mapper) MapPSQLRowsToLeaderboardData(rows *sql.Rows) (leaderboardData models.PSQLLeaderboardDataList) {
	var data models.PSQLLeaderboardData
	for rows.Next() {
		err := rows.Scan(
			&data.CompId,
			&data.CompName,
			&data.TeamId,
			&data.TeamName,
			&data.GamesPlayed,
			&data.WinCount,
			&data.DrawCount,
			&data.LossCount,
			&data.Bye,
			&data.PointsFor,
			&data.PointsAgainst,
			&data.TriesFor,
			&data.TriesAgainst,
			&data.BonusPointsTry,
			&data.BonusPointsLosing,
			&data.BonusPoints,
			&data.PointsDiff,
			&data.Points,
		)
		if err != nil {
			log.Panicln(err)
		}
		leaderboardData = append(leaderboardData, data)
	}

	return leaderboardData
}

func (m Mapper) MapPSQLLeaderboardDataToResponse(compId, compName string, leaderboardData models.PSQLLeaderboardDataList) (resp response.LeaderboardResponse) {
	resp.Id = compId
	resp.Name = compName
	for _, data := range leaderboardData {
		resp.Teams = append(resp.Teams, models.TeamLeaderboardData{
			Id:                strconv.Itoa(data.TeamId),
			Name:              data.TeamName,
			GamesPlayed:       data.GamesPlayed,
			WinCount:          data.WinCount,
			DrawCount:         data.DrawCount,
			LossCount:         data.LossCount,
			Bye:               data.Bye,
			PointsFor:         data.PointsFor,
			PointsAgainst:     data.PointsAgainst,
			TriesFor:          data.TriesFor,
			TriesAgainst:      data.TriesAgainst,
			BonusPointsTry:    data.BonusPointsTry,
			BonusPointsLosing: data.BonusPointsLosing,
			BonusPoints:       data.BonusPoints,
			PointsDiff:        data.PointsDiff,
			Points:            data.Points,
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
			c.comp_id = '%v'
	) select comp_id, comp_name, team_id, team_name from comp_teams;`

	PSQLLeaderboardById = `select * from public.comp_leaderboard where comp_id = '%v'`
)
