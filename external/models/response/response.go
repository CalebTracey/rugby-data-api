package response

import "github.com/calebtracey/rugby-data-api/external/models"

type TeamDataResponse struct {
	Teams models.PSQLTeamDataList `json:"teams"`
}

type CompetitionResponse struct {
	Teams   models.PSQLCompetitionDataList `json:"teams,omitempty"`
	Message Message                        `json:"message,omitempty"`
}

type PSQLResponse struct {
	RowsAffected string  `json:"rowsAffected,omitempty"`
	LastInsertID string  `json:"lastInsertID,omitempty"`
	Message      Message `json:"message,omitempty"`
}

type Message struct {
	ErrorLog  []ErrorLog `json:"errorLog,omitempty"`
	HostName  string     `json:"hostName,omitempty"`
	Status    string     `json:"status,omitempty"`
	TimeTaken string     `json:"timeTaken,omitempty"`
	Count     int        `json:"count,omitempty"`
}

type ErrorLogs []ErrorLog

type ErrorLog struct {
	Scope      string `json:"scope,omitempty"`
	StatusCode string `json:"status,omitempty"`
	Trace      string `json:"trace,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
	Query      string `json:"query,omitempty"`
}
