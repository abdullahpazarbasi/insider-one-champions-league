package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestComputeStandingsCalculatesPoints(t *testing.T) {
	calculator := TableStandingsCalculator{}
	scoreHome := 2
	scoreAway := 1
	weeks := []domain.Week{{Matches: []domain.Match{{HomeTeamID: "a", AwayTeamID: "b", HomeScore: &scoreHome, AwayScore: &scoreAway}}}}
	teams := []domain.Team{{ID: "a", Name: "Team A"}, {ID: "b", Name: "Team B"}}

	standings := calculator.ComputeStandings(weeks, teams)

	if len(standings) != 2 {
		t.Fatalf("expected standings for each team")
	}

	if standings[0].Points != 3 || standings[0].Wins != 1 {
		t.Fatalf("expected winner to get three points: %+v", standings[0])
	}
}

func TestComputeStandingsSkipsUnknownTeams(t *testing.T) {
	calculator := TableStandingsCalculator{}
	scoreHome := 0
	scoreAway := 0
	weeks := []domain.Week{{Matches: []domain.Match{{HomeTeamID: "unknown", AwayTeamID: "b", HomeScore: &scoreHome, AwayScore: &scoreAway}}}}
	teams := []domain.Team{{ID: "b", Name: "Team B"}}

	standings := calculator.ComputeStandings(weeks, teams)

	if len(standings) != 1 {
		t.Fatalf("expected only known teams to be included")
	}
}
