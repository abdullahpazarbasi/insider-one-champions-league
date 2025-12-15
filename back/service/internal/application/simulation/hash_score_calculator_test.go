package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestCalculateScoreIsDeterministic(t *testing.T) {
	calculator := HashScoreCalculator{}
	home := domain.Team{ID: "h", Strength: 90}
	away := domain.Team{ID: "a", Strength: 80}

	firstHome, firstAway := calculator.CalculateScore(home, away, 1, 1)
	secondHome, secondAway := calculator.CalculateScore(home, away, 1, 1)

	if firstHome != secondHome || firstAway != secondAway {
		t.Fatalf("expected deterministic scores, got %d-%d and %d-%d", firstHome, firstAway, secondHome, secondAway)
	}
}

func TestClampScoreLimitsValues(t *testing.T) {
	if clampScore(-5) != 0 || clampScore(6) != 5 {
		t.Fatalf("clampScore failed to limit values")
	}
}

func TestCalculateScoreAccountsForBias(t *testing.T) {
	calculator := HashScoreCalculator{}
	strongHome := domain.Team{ID: "home-strong", Strength: 95}
	weakAway := domain.Team{ID: "away-weak", Strength: 40}
	weakHome := domain.Team{ID: "home-weak", Strength: 40}
	strongAway := domain.Team{ID: "away-strong", Strength: 95}

	boostedHome, _ := calculator.CalculateScore(strongHome, weakAway, 2, 2)
	limitedHome, _ := calculator.CalculateScore(weakHome, strongAway, 2, 2)

	if boostedHome <= limitedHome {
		t.Fatalf("expected strength bias to favor stronger home team")
	}
}
