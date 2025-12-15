package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type StandingsCalculator interface {
	ComputeStandings(weeks []domain.Week, teams []domain.Team) []domain.Standing
}
