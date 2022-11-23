package models

type TeamData struct {
	Name    string   `json:"name,omitempty"`
	Players []Player `json:"players,omitempty"`
}

type Player struct {
	Name     string `json:"name,omitempty"`
	Position string `json:"position,omitempty"`
}
