package simulation

import (
	"fmt"
	"strings"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (StrictTeamValidator) ValidateTeams(teams []domain.Team) error {
	if len(teams) == 0 {
		return fmt.Errorf("at least one team is required")
	}

	for _, t := range teams {
		if strings.TrimSpace(t.Name) == "" {
			return fmt.Errorf("team %s is missing a name", t.ID)
		}
		if t.Strength < 0 || t.Strength > 100 {
			return fmt.Errorf("team %s strength must be between 0 and 100", t.Name)
		}
	}

	return nil
}
