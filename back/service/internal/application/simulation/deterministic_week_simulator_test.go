package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestApplySimulationRespectsManualScores(t *testing.T) {
	simulator := NewDeterministicWeekSimulator(HashScoreCalculator{})
	score := 2
	weeks := []domain.Week{{WeekNumber: 1, Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b", HomeScore: &score, AwayScore: &score}}}}
	teams := []domain.Team{{ID: "a", Strength: 90}, {ID: "b", Strength: 80}}

	simulated := simulator.ApplySimulation(weeks, teams, 0)

	if simulated[0].Matches[0].HomeScore == nil || *simulated[0].Matches[0].HomeScore != 2 {
		t.Fatalf("manual score overwritten: %+v", simulated[0].Matches[0])
	}
}

func TestApplySimulationIgnoresNegativeTarget(t *testing.T) {
	simulator := NewDeterministicWeekSimulator(HashScoreCalculator{})
	weeks := []domain.Week{{WeekNumber: 1, Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b"}}}}
	teams := []domain.Team{{ID: "a", Strength: 90}, {ID: "b", Strength: 80}}

	simulated := simulator.ApplySimulation(weeks, teams, -1)
	if simulated[0].Matches[0].HomeScore != nil {
		t.Fatalf("expected no simulation for negative target")
	}
}

func TestApplySimulationCapsTargetToExistingWeeks(t *testing.T) {
	simulator := NewDeterministicWeekSimulator(HashScoreCalculator{})
	weeks := []domain.Week{{WeekNumber: 1, Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b"}}}}
	teams := []domain.Team{{ID: "a", Strength: 90}, {ID: "b", Strength: 80}}

	simulated := simulator.ApplySimulation(weeks, teams, 5)
	if simulated[0].Matches[0].HomeScore == nil {
		t.Fatalf("expected first week to be simulated when target exceeds length")
	}
}

func TestFillWeekSkipsUnknownTeams(t *testing.T) {
	simulator := NewDeterministicWeekSimulator(HashScoreCalculator{})
	weeks := []domain.Week{{WeekNumber: 1, Matches: []domain.Match{{HomeTeamID: "missing", AwayTeamID: "b"}}}}
	teams := []domain.Team{{ID: "a", Strength: 90}, {ID: "b", Strength: 80}}

	simulated := simulator.ApplySimulation(weeks, teams, 0)
	if simulated[0].Matches[0].HomeScore != nil || simulated[0].Matches[0].AwayScore != nil {
		t.Fatalf("expected missing teams to be ignored")
	}
}
