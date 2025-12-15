package simulation

func (TwoDecimalRounder) RoundToDisplay(value float64) float64 {
	scaled := value * 100
	truncated := float64(int(scaled + 0.5))
	return truncated / 100
}
