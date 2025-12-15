package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type TeamValidator interface {
	ValidateTeams(teams []domain.Team) error
}
