package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func NewStaticTeamProvider() StaticTeamProvider {
	return StaticTeamProvider{teams: []domain.Team{
		{ID: "team-1", Name: "Real Madrid", Strength: 95},
		{ID: "team-2", Name: "Bayern Munich", Strength: 92},
		{ID: "team-3", Name: "Manchester City", Strength: 94},
		{ID: "team-4", Name: "Inter", Strength: 89},
	}}
}
