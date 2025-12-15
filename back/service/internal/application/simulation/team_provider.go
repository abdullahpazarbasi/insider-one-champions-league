package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type TeamProvider interface {
	DefaultTeams() []domain.Team
}
