package domain

type Week struct {
	WeekNumber int     `json:"weekNumber"`
	Matches    []Match `json:"matches"`
}
