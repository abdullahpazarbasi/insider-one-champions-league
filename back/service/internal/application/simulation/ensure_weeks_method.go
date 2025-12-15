package simulation

import (
	"fmt"
	"sort"
	"strings"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (s WeekBuilderService) EnsureWeeks(weeks []domain.Week, teams []domain.Team) []domain.Week {
	if len(weeks) == 0 {
		return s.scheduleGenerator.GenerateSchedule(teams)
	}

	for i := range weeks {
		if weeks[i].WeekNumber == 0 {
			weeks[i].WeekNumber = i + 1
		}
		for m := range weeks[i].Matches {
			if strings.TrimSpace(weeks[i].Matches[m].ID) == "" {
				weeks[i].Matches[m].ID = fmt.Sprintf("w%d-m%d", weeks[i].WeekNumber, m+1)
			}
		}
	}

	sort.SliceStable(weeks, func(i, j int) bool {
		return weeks[i].WeekNumber < weeks[j].WeekNumber
	})

	return weeks
}
