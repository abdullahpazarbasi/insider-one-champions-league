package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type TeamIDAssigner interface {
	EnsureTeamIDs(teams []domain.Team) []domain.Team
}
