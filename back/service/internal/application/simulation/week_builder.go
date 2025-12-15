package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type WeekBuilder interface {
	EnsureWeeks(weeks []domain.Week, teams []domain.Team) []domain.Week
}
