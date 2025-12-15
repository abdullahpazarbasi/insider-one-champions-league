package simulation

import "testing"

func TestRoundToDisplayRoundsHalfUp(t *testing.T) {
	rounder := TwoDecimalRounder{}

	if rounder.RoundToDisplay(1.234) != 1.23 {
		t.Fatalf("expected rounding down")
	}
	if rounder.RoundToDisplay(1.235) != 1.24 {
		t.Fatalf("expected rounding up")
	}
}
