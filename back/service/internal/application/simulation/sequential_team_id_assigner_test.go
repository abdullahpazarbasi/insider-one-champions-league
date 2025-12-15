package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestEnsureTeamIDsFillsEmptyValues(t *testing.T) {
	assigner := SequentialTeamIDAssigner{}
	teams := []domain.Team{{Name: "Team 1"}, {Name: "Team 2"}}

	updated := assigner.EnsureTeamIDs(teams)

	if updated[0].ID == "" || updated[1].ID == "" {
		t.Fatalf("expected IDs to be assigned: %+v", updated)
	}
}
