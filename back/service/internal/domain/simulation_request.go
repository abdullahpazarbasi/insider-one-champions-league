package domain

type SimulationRequest struct {
	Teams           []Team `json:"teams"`
	Weeks           []Week `json:"weeks"`
	TargetWeekIndex *int   `json:"targetWeekIndex,omitempty"`
}
