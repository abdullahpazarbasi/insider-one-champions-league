package simulation

import (
	"sort"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (TableStandingsCalculator) ComputeStandings(weeks []domain.Week, teams []domain.Team) []domain.Standing {
	stats := make(map[string]*domain.Standing)
	for _, t := range teams {
		stats[t.ID] = &domain.Standing{TeamID: t.ID, Name: t.Name}
	}

	for _, w := range weeks {
		for _, m := range w.Matches {
			if m.HomeScore == nil || m.AwayScore == nil {
				continue
			}

			home := stats[m.HomeTeamID]
			away := stats[m.AwayTeamID]
			if home == nil || away == nil {
				continue
			}

			home.Played++
			away.Played++

			home.GoalsFor += *m.HomeScore
			home.GoalsAgainst += *m.AwayScore
			away.GoalsFor += *m.AwayScore
			away.GoalsAgainst += *m.HomeScore

			switch {
			case *m.HomeScore > *m.AwayScore:
				home.Wins++
				away.Losses++
				home.Points += 3
			case *m.HomeScore < *m.AwayScore:
				away.Wins++
				home.Losses++
				away.Points += 3
			default:
				home.Draws++
				away.Draws++
				home.Points++
				away.Points++
			}
		}
	}

	standings := []domain.Standing{}
	for _, s := range stats {
		s.GoalDifference = s.GoalsFor - s.GoalsAgainst
		standings = append(standings, *s)
	}

	sort.SliceStable(standings, func(i, j int) bool {
		if standings[i].Points == standings[j].Points {
			return standings[i].GoalDifference > standings[j].GoalDifference
		}
		return standings[i].Points > standings[j].Points
	})

	return standings
}
