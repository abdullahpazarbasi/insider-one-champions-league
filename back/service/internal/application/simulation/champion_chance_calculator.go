package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type ChampionChanceCalculator interface {
	CalculateChances(standings []domain.Standing, teams []domain.Team) []domain.ChampionChance
}
