package simulation

func NewWeightedChampionChanceCalculator(rounder PercentageRounder) WeightedChampionChanceCalculator {
	return WeightedChampionChanceCalculator{rounder: rounder}
}
