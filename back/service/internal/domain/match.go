package domain

type Match struct {
	ID         string `json:"id"`
	HomeTeamID string `json:"homeTeamId"`
	AwayTeamID string `json:"awayTeamId"`
	HomeScore  *int   `json:"homeScore,omitempty"`
	AwayScore  *int   `json:"awayScore,omitempty"`
}
