package models

// PSQLCompetitionDataList struct
//
//swagger:model
type PSQLCompetitionDataList []PSQLCompetitionData

// PSQLCompetitionData struct
//
//swagger:model
type PSQLCompetitionData struct {
	CompID   int    `json:"comp_id,omitempty"`
	CompName string `json:"compName,omitempty"`
	TeamID   int    `json:"team_id,omitempty"`
	TeamName string `json:"team_name,omitempty"`
}

// TeamDataList struct
//
//swagger:model
type TeamDataList []TeamData

// TeamData struct
//
//swagger:model
type TeamData struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// PSQLTeamDataList struct
//
//swagger:model
type PSQLTeamDataList []PSQLTeamData

// PSQLTeamData struct
//
//swagger:model
type PSQLTeamData struct {
	TeamID int    `json:"team_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Player struct {
	Name     string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
}
