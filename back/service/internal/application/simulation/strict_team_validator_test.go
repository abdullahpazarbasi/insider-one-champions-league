package simulation

import (
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func TestValidateTeamsRejectsEmpty(t *testing.T) {
	validator := StrictTeamValidator{}
	if err := validator.ValidateTeams(nil); err == nil {
		t.Fatalf("expected error for empty teams")
	}
}

func TestValidateTeamsRejectsInvalidStrength(t *testing.T) {
	validator := StrictTeamValidator{}
	teams := []domain.Team{{ID: "t1", Name: "Team 1", Strength: 200}}

	if err := validator.ValidateTeams(teams); err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestValidateTeamsRejectsMissingName(t *testing.T) {
	validator := StrictTeamValidator{}
	teams := []domain.Team{{ID: "t1", Name: "   ", Strength: 10}}

	if err := validator.ValidateTeams(teams); err == nil {
		t.Fatalf("expected validation error for empty name")
	}
}

func TestValidateTeamsAcceptsValid(t *testing.T) {
	validator := StrictTeamValidator{}
	teams := []domain.Team{{ID: "t1", Name: "Team 1", Strength: 50}}

	if err := validator.ValidateTeams(teams); err != nil {
		t.Fatalf("expected no validation error, got %v", err)
	}
}
