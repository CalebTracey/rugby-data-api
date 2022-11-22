package request

type CompetitionRequest struct {
	Source          string `json:"source,omitempty"`
	CompetitionID   string `json:"competitionID,omitempty"`
	CompetitionName string `json:"competitionName,omitempty"`
	Table           string `json:"table,omitempty"`
	ID              string `json:"ID,omitempty"`
}
