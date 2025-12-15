package simulation

type PercentageRounder interface {
	RoundToDisplay(value float64) float64
}
