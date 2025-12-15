package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type WeekSimulator interface {
	ApplySimulation(weeks []domain.Week, teams []domain.Team, targetWeekIndex int) []domain.Week
}
