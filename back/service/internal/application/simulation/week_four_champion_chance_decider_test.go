package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestShouldCalculateChampionChancesRequiresFourthWeekComplete(t *testing.T) {
	checker := MatchScoreWeekCompletionChecker{}
	decider := NewWeekFourChampionChanceDecider(checker)
	week := domain.Week{WeekNumber: 4, Matches: []domain.Match{{}}}

	if decider.ShouldCalculateChampionChances([]domain.Week{week}) {
		t.Fatalf("expected false when scores missing")
	}

	score := 1
	week.Matches[0].HomeScore = &score
	week.Matches[0].AwayScore = &score
	if !decider.ShouldCalculateChampionChances([]domain.Week{week}) {
		t.Fatalf("expected true once week four is complete")
	}
}

func TestShouldCalculateChampionChancesSkipsWhenWeekNotPresent(t *testing.T) {
	checker := MatchScoreWeekCompletionChecker{}
	decider := NewWeekFourChampionChanceDecider(checker)

	if decider.ShouldCalculateChampionChances([]domain.Week{{WeekNumber: 1}}) {
		t.Fatalf("expected false when week four is absent")
	}
}
