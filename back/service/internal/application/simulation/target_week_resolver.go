package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type TargetWeekResolver interface {
	ResolveTargetWeekIndex(targetWeekIndex *int, weeks []domain.Week) (int, error)
}
