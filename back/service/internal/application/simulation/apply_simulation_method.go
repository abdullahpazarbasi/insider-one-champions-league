package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (s DeterministicWeekSimulator) ApplySimulation(weeks []domain.Week, teams []domain.Team, targetWeekIndex int) []domain.Week {
	if targetWeekIndex < 0 {
		return weeks
	}

	teamMap := make(map[string]domain.Team)
	for _, t := range teams {
		teamMap[t.ID] = t
	}

	if targetWeekIndex >= len(weeks) {
		targetWeekIndex = len(weeks) - 1
	}

	for weekIdx := 0; weekIdx <= targetWeekIndex; weekIdx++ {
		fillWeek(&weeks[weekIdx], teamMap, s.scoreCalculator)
	}

	return weeks
}
