package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestIsWeekCompleteRequiresAllScores(t *testing.T) {
	checker := MatchScoreWeekCompletionChecker{}
	score := 1
	week := domain.Week{WeekNumber: 1, Matches: []domain.Match{{HomeScore: &score, AwayScore: &score}, {HomeScore: nil, AwayScore: nil}}}

	if checker.IsWeekComplete(week) {
		t.Fatalf("expected incomplete week when scores are missing")
	}
}
