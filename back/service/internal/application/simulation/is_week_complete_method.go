package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (MatchScoreWeekCompletionChecker) IsWeekComplete(week domain.Week) bool {
	for _, m := range week.Matches {
		if m.HomeScore == nil || m.AwayScore == nil {
			return false
		}
	}

	return len(week.Matches) > 0
}
