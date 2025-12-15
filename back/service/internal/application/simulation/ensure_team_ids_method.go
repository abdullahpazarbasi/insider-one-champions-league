package simulation

import (
	"fmt"
	"strings"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (SequentialTeamIDAssigner) EnsureTeamIDs(teams []domain.Team) []domain.Team {
	for i := range teams {
		if strings.TrimSpace(teams[i].ID) == "" {
			teams[i].ID = fmt.Sprintf("team-%d", i+1)
		}
	}
	return teams
}
