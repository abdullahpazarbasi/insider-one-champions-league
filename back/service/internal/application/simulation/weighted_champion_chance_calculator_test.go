package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestCalculateChancesUsesRounder(t *testing.T) {
	rounder := TwoDecimalRounder{}
	calculator := NewWeightedChampionChanceCalculator(rounder)
	standings := []domain.Standing{{TeamID: "a", Points: 10, GoalDifference: 5}, {TeamID: "b", Points: 5, GoalDifference: 0}}
	teams := []domain.Team{{ID: "a", Strength: 90}, {ID: "b", Strength: 80}}

	chances := calculator.CalculateChances(standings, teams)

	if len(chances) != 2 {
		t.Fatalf("expected two chance entries")
	}

	if chances[0].Percentage <= chances[1].Percentage {
		t.Fatalf("expected stronger record to have higher chance")
	}
}

func TestCalculateChancesFloorsWeights(t *testing.T) {
	rounder := TwoDecimalRounder{}
	calculator := NewWeightedChampionChanceCalculator(rounder)
	standings := []domain.Standing{{TeamID: "a", Points: 0, GoalDifference: -100}}
	teams := []domain.Team{{ID: "a", Strength: 0}}

	chances := calculator.CalculateChances(standings, teams)
	if chances[0].Percentage != 100 {
		t.Fatalf("expected single team chance to sum to 100 with floored weight")
	}
}
