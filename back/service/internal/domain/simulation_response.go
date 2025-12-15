package domain

type SimulationResponse struct {
	Teams           []Team           `json:"teams"`
	Weeks           []Week           `json:"weeks"`
	Standings       []Standing       `json:"standings"`
	ChampionChances []ChampionChance `json:"championChances,omitempty"`
}
