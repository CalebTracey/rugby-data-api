package response

import "github.com/calebtracey/rugby-data-api/external/models"

// TeamDataResponse struct
//
//swagger:model
type TeamDataResponse struct {
	Teams models.PSQLTeamDataList `json:"teams"`
}

// CompetitionResponse struct
//
//swagger:model
type CompetitionResponse struct {
	ID      string              `json:"id,omitempty"`
	Name    string              `json:"name,omitempty"`
	Teams   models.TeamDataList `json:"teams,omitempty"`
	Message Message             `json:"message,omitempty"`
}

// CompetitionCrawlResponse struct
//
//swagger:model
type CompetitionCrawlResponse struct {
	CompetitionID string  `json:"id,omitempty"`
	Message       Message `json:"message,omitempty"`
}

// Message struct
//
//swagger:model
type Message struct {
	ErrorLog  ErrorLogs `json:"errorLog,omitempty"`
	HostName  string    `json:"hostName,omitempty"`
	Status    string    `json:"status,omitempty"`
	TimeTaken string    `json:"timeTaken,omitempty"`
	Count     int       `json:"count,omitempty"`
}

// ErrorLogs struct
//
//swagger:model
type ErrorLogs []ErrorLog

// ErrorLog struct
//
//swagger:model
type ErrorLog struct {
	Scope      string `json:"scope,omitempty"`
	StatusCode string `json:"status,omitempty"`
	Trace      string `json:"trace,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
	Query      string `json:"query,omitempty"`
}
