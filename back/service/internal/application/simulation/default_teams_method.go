package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (p StaticTeamProvider) DefaultTeams() []domain.Team {
	return append([]domain.Team{}, p.teams...)
}
