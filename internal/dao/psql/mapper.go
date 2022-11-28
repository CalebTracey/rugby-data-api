package psql

import (
	"database/sql"
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/calebtracey/rugby-models/pkg/models"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//go:generate mockgen -destination=../../mocks/dbmocks/mockMapper.go -package=dbmocks . MapperI
type MapperI interface {
	CreatePSQLLeaderboardByIdQuery(teamId string) string
	MapPSQLRowsToLeaderboardData(rows *sql.Rows) (leaderboardData models.PSQLLeaderboardDataList, errorLog *response.ErrorLog)
	MapPSQLLeaderboardDataToResponse(compId, compName string, leaderboardData models.PSQLLeaderboardDataList) (resp dtos.CompetitionLeaderboardData)
	MapPSQLAllLeaderboardDataToResponse(leaderboardDataList models.PSQLLeaderboardDataList) (resp response.LeaderboardResponse)
}

type Mapper struct{}

func (m Mapper) CreatePSQLLeaderboardByIdQuery(teamId string) string {
	teamIdInt, err := strconv.Atoi(teamId)
	if err != nil {
		log.Error(err)
		return ""
	}
	return fmt.Sprintf(LeaderboardByIdQuery, teamIdInt)
}

func (m Mapper) MapPSQLRowsToLeaderboardData(rows *sql.Rows) (leaderboardData models.PSQLLeaderboardDataList, errorLog *response.ErrorLog) {
	var data models.PSQLLeaderboardData
	for rows.Next() {
		if err := rows.Scan(
			&data.CompId,
			&data.CompName,
			&data.TeamId,
			&data.TeamName,
			&data.TeamAbbr,
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
		); err != nil {
			return leaderboardData, mapError(err, "")
		}

		leaderboardData = append(leaderboardData, data)
	}

	return leaderboardData, nil
}

func mapPSQLTeamData(data models.PSQLLeaderboardData) dtos.TeamLeaderboardData {
	return dtos.TeamLeaderboardData{
		Id:   strconv.Itoa(data.TeamId),
		Name: data.TeamName,
		Abbr: data.TeamAbbr,
		CompetitionStats: dtos.TeamCompetitionStats{
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
		},
	}
}

func (m Mapper) MapPSQLLeaderboardDataToResponse(compId, compName string, leaderboardData models.PSQLLeaderboardDataList) (resp dtos.CompetitionLeaderboardData) {
	resp.CompName = compName
	resp.CompId = compId

	for _, data := range leaderboardData {
		resp.Teams = append(resp.Teams, mapPSQLTeamData(data))
	}
	return resp
}

func (m Mapper) MapPSQLAllLeaderboardDataToResponse(leaderboardDataList models.PSQLLeaderboardDataList) (resp response.LeaderboardResponse) {
	compDataMap := make(map[string]dtos.CompetitionLeaderboardData, CompetitionCount)
	for _, data := range leaderboardDataList {
		if comp, ok := compDataMap[data.CompName]; ok {
			comp.Teams = append(comp.Teams, mapPSQLTeamData(data))
			compDataMap[data.CompName] = comp
		} else {
			comp.CompId = strconv.Itoa(data.CompId)
			comp.CompName = data.CompName
			comp.Teams = dtos.TeamLeaderboardDataList{mapPSQLTeamData(data)}
			compDataMap[data.CompName] = comp
		}
	}
	for _, team := range compDataMap {
		resp.LeaderboardData = append(resp.LeaderboardData, team)
	}
	return resp
}

const (
	CompetitionCount = 6

	CompetitionByID = `with comp_teams as (
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

	LeaderboardByIdQuery = `select * from public.comp_leaderboard where comp_id = '%v'`

	AllLeaderboardsQuery = `select * from public.comp_leaderboard`
)
