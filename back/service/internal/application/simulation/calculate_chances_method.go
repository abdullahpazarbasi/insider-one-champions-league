package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (c WeightedChampionChanceCalculator) CalculateChances(standings []domain.Standing, teams []domain.Team) []domain.ChampionChance {
	teamStrength := make(map[string]int)
	for _, t := range teams {
		teamStrength[t.ID] = t.Strength
	}

	weights := make([]float64, len(standings))
	total := 0.0
	for i, s := range standings {
		strength := float64(teamStrength[s.TeamID])
		weight := float64(s.Points*4+s.GoalDifference*2) + strength
		if weight < 1 {
			weight = 1
		}
		weights[i] = weight
		total += weight
	}

	chances := make([]domain.ChampionChance, len(standings))
	for i, s := range standings {
		percentage := (weights[i] / total) * 100
		chances[i] = domain.ChampionChance{TeamID: s.TeamID, Percentage: c.rounder.RoundToDisplay(percentage)}
	}

	return chances
}
