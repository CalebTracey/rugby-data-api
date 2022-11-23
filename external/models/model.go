package models

type PSQLCompetitionDataList []PSQLCompetitionData

type PSQLCompetitionData struct {
	CompID   int    `json:"comp_id,omitempty"`
	CompName string `json:"compName,omitempty"`
	TeamID   int    `json:"team_id,omitempty"`
	TeamName string `json:"team_name,omitempty"`
}

type TeamDataList []TeamData

type TeamData struct {
	ID   string
	Name string
}

type PSQLTeamDataList []PSQLTeamData

type PSQLTeamData struct {
	TeamID int    `json:"team_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Player struct {
	Name     string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
}
