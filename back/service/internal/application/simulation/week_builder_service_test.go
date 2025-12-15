package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestEnsureWeeksGeneratesWhenEmpty(t *testing.T) {
	generator := RoundRobinScheduleGenerator{}
	builder := NewWeekBuilderService(generator)
	teams := []domain.Team{{ID: "a", Name: "A"}, {ID: "b", Name: "B"}}

	weeks := builder.EnsureWeeks(nil, teams)

	if len(weeks) == 0 {
		t.Fatalf("expected schedule to be generated")
	}
}

func TestEnsureWeeksAddsIdentifiers(t *testing.T) {
	generator := RoundRobinScheduleGenerator{}
	builder := NewWeekBuilderService(generator)
	weeks := []domain.Week{{Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b"}}}}

	enriched := builder.EnsureWeeks(weeks, []domain.Team{{ID: "a"}, {ID: "b"}})

	if enriched[0].WeekNumber == 0 || enriched[0].Matches[0].ID == "" {
		t.Fatalf("expected identifiers to be assigned: %+v", enriched)
	}
}
