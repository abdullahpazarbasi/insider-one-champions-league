package simulation

import (
	"fmt"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (TargetWeekResolverService) ResolveTargetWeekIndex(targetWeekIndex *int, weeks []domain.Week) (int, error) {
	if targetWeekIndex == nil || len(weeks) == 0 {
		return -1, nil
	}

	idx := *targetWeekIndex
	if idx < 0 {
		return -1, fmt.Errorf("targetWeekIndex cannot be negative")
	}

	if idx >= len(weeks) {
		return len(weeks) - 1, nil
	}

	return idx, nil
}
