package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type ScoreCalculator interface {
	CalculateScore(home domain.Team, away domain.Team, weekNumber int, matchIndex int) (int, int)
}
