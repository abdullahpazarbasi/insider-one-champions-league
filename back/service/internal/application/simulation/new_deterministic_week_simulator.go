package simulation

func NewDeterministicWeekSimulator(calculator ScoreCalculator) DeterministicWeekSimulator {
	return DeterministicWeekSimulator{scoreCalculator: calculator}
}
