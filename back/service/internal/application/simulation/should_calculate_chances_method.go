package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (d WeekFourChampionChanceDecider) ShouldCalculateChampionChances(weeks []domain.Week) bool {
	for _, w := range weeks {
		if w.WeekNumber == 4 {
			return d.weekChecker.IsWeekComplete(w)
		}
	}

	return false
}
