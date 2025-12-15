package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type WeekCompletionChecker interface {
	IsWeekComplete(week domain.Week) bool
}
