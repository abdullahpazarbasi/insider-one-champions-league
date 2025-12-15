package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (s SeasonSimulationService) resolveChampionChances(weeks []domain.Week, standings []domain.Standing, teams []domain.Team) []domain.ChampionChance {
	if !s.championChanceDecider.ShouldCalculateChampionChances(weeks) {
		return nil
	}

	return s.championChanceCalculator.CalculateChances(standings, teams)
}
