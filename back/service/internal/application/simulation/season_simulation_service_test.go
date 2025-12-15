package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestBootstrapSimulatesOpeningWeek(t *testing.T) {
	service := NewDefaultSimulationService()

	response := service.Bootstrap()

	if len(response.Weeks) == 0 || len(response.Weeks[0].Matches) == 0 {
		t.Fatalf("bootstrap returned empty schedule")
	}

	if response.Weeks[0].Matches[0].HomeScore == nil {
		t.Fatalf("expected first week to be simulated")
	}

	if response.ChampionChances != nil {
		t.Fatalf("champion chances should be omitted until week four")
	}
}

func TestSimulateUsesTargetWeek(t *testing.T) {
	service := NewDefaultSimulationService()
	teams := []domain.Team{{Name: "A", Strength: 90}, {Name: "B", Strength: 80}, {Name: "C", Strength: 85}, {Name: "D", Strength: 70}}
	target := 1

	response, err := service.Simulate(domain.SimulationRequest{Teams: teams, TargetWeekIndex: &target})
	if err != nil {
		t.Fatalf("simulate returned error: %v", err)
	}

	if len(response.Weeks) < 2 || response.Weeks[1].Matches[0].HomeScore == nil {
		t.Fatalf("expected simulation to fill weeks up to target")
	}
}

func TestSimulateHonorsManualScoresAndProducesChances(t *testing.T) {
	service := NewDefaultSimulationService()
	teams := []domain.Team{{ID: "a", Name: "A", Strength: 90}, {ID: "b", Name: "B", Strength: 80}, {ID: "c", Name: "C", Strength: 85}, {ID: "d", Name: "D", Strength: 70}}
	weeks := NewWeekBuilderService(RoundRobinScheduleGenerator{}).EnsureWeeks(nil, teams)
	score := 3
	weeks[0].Matches[0].HomeScore = &score
	weeks[0].Matches[0].AwayScore = &score
	target := len(weeks) - 1

	response, err := service.Simulate(domain.SimulationRequest{Teams: teams, Weeks: weeks, TargetWeekIndex: &target})
	if err != nil {
		t.Fatalf("simulate returned error: %v", err)
	}

	if response.Weeks[0].Matches[0].HomeScore == nil || *response.Weeks[0].Matches[0].HomeScore != 3 {
		t.Fatalf("manual score overwritten: %+v", response.Weeks[0].Matches[0])
	}

	if len(response.ChampionChances) != len(teams) {
		t.Fatalf("expected champion chances once week four completed")
	}
}

func TestSimulateFailsValidation(t *testing.T) {
	service := NewDefaultSimulationService()
	_, err := service.Simulate(domain.SimulationRequest{Teams: []domain.Team{{Name: " ", Strength: 10}}})

	if err == nil {
		t.Fatalf("expected validation error for invalid teams")
	}
}

func TestSimulateSkipsWhenTargetMissing(t *testing.T) {
	service := NewDefaultSimulationService()
	teams := []domain.Team{{ID: "a", Name: "A", Strength: 90}, {ID: "b", Name: "B", Strength: 80}}
	weeks := []domain.Week{{WeekNumber: 1, Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b"}}}}

	response, err := service.Simulate(domain.SimulationRequest{Teams: teams, Weeks: weeks})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response.Weeks[0].Matches[0].HomeScore != nil {
		t.Fatalf("expected simulation to skip when target week is absent")
	}
}
