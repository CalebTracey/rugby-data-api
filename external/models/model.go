package models

type PSQLCompetitionDataList []PSQLCompetitionData

type PSQLCompetitionData struct {
	CompID   int    `json:"comp_id,omitempty"`
	CompName string `json:"compName,omitempty"`
	TeamID   int    `json:"team_id,omitempty"`
	TeamName string `json:"team_name,omitempty"`
	//Teams  PSQLTeamDataList `json:"teams,omitempty"`
}

type PSQLTeamDataList []PSQLTeamData

type PSQLTeamData struct {
	TeamID int    `json:"team_id,omitempty"`
	CompID int    `json:"comp_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Player struct {
	Name     string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
}
