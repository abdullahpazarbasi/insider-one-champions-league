package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestResolveTargetWeekIndexDefaultsToLast(t *testing.T) {
	resolver := TargetWeekResolverService{}
	weeks := make([]domain.Week, 2)
	idx := 5

	resolved, err := resolver.ResolveTargetWeekIndex(&idx, weeks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resolved != 1 {
		t.Fatalf("expected last index when provided index is large, got %d", resolved)
	}
}

func TestResolveTargetWeekIndexRejectsNegative(t *testing.T) {
	resolver := TargetWeekResolverService{}
	weeks := make([]domain.Week, 1)
	idx := -1

	_, err := resolver.ResolveTargetWeekIndex(&idx, weeks)
	if err == nil {
		t.Fatalf("expected error for negative index")
	}
}

func TestResolveTargetWeekIndexAllowsNil(t *testing.T) {
	resolver := TargetWeekResolverService{}
	weeks := make([]domain.Week, 1)

	resolved, err := resolver.ResolveTargetWeekIndex(nil, weeks)
	if err != nil || resolved != -1 {
		t.Fatalf("expected nil target to default to -1: %d %v", resolved, err)
	}
}
