package simulation

// DeterministicWeekSimulator fills missing scores using deterministic rules.
type DeterministicWeekSimulator struct {
	scoreCalculator ScoreCalculator
}
