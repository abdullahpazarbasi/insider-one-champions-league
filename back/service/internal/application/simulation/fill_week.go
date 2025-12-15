package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func fillWeek(w *domain.Week, teamMap map[string]domain.Team, calculator ScoreCalculator) {
	for matchIdx := range w.Matches {
		m := &w.Matches[matchIdx]
		if m.HomeScore != nil && m.AwayScore != nil {
			continue
		}

		home, okHome := teamMap[m.HomeTeamID]
		away, okAway := teamMap[m.AwayTeamID]
		if !okHome || !okAway {
			continue
		}

		homeScore, awayScore := calculator.CalculateScore(home, away, w.WeekNumber, matchIdx)
		m.HomeScore = &homeScore
		m.AwayScore = &awayScore
	}
}
