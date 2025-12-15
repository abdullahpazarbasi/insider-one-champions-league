package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type ChampionChanceDecider interface {
	ShouldCalculateChampionChances(weeks []domain.Week) bool
}
