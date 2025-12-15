package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestGenerateScheduleProducesHomeAndAwayRounds(t *testing.T) {
	generator := RoundRobinScheduleGenerator{}
	teams := []domain.Team{{ID: "a"}, {ID: "b"}, {ID: "c"}, {ID: "d"}}

	weeks := generator.GenerateSchedule(teams)

	expectedWeeks := (len(teams) - 1) * 2
	if len(weeks) != expectedWeeks {
		t.Fatalf("expected %d weeks, got %d", expectedWeeks, len(weeks))
	}

	if weeks[0].Matches[0].HomeTeamID == weeks[len(weeks)-1].Matches[0].HomeTeamID {
		t.Fatalf("expected second leg to swap home and away teams")
	}
}

func TestGenerateScheduleHandlesInsufficientTeams(t *testing.T) {
	generator := RoundRobinScheduleGenerator{}
	weeks := generator.GenerateSchedule([]domain.Team{{ID: "only"}})

	if len(weeks) != 0 {
		t.Fatalf("expected no schedule for fewer than two teams")
	}
}
