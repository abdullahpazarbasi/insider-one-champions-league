package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (RoundRobinScheduleGenerator) GenerateSchedule(teams []domain.Team) []domain.Week {
	schedule := []domain.Week{}
	if len(teams) < 2 {
		return schedule
	}

	ids := make([]string, len(teams))
	for i, t := range teams {
		ids[i] = t.ID
	}

	rotations := ids[1:]
	rounds := len(teams) - 1

	for round := 0; round < rounds; round++ {
		matches := []domain.Match{}
		first := ids[0]
		opponent := rotations[round%len(rotations)]
		matches = append(matches, domain.Match{ID: matchID(round+1, 1), HomeTeamID: first, AwayTeamID: opponent})

		for i := 1; i < len(teams)/2; i++ {
			home := rotations[(round+i)%len(rotations)]
			away := rotations[(round+len(rotations)-i)%len(rotations)]
			matches = append(matches, domain.Match{ID: matchID(round+1, i+1), HomeTeamID: home, AwayTeamID: away})
		}

		schedule = append(schedule, domain.Week{WeekNumber: round + 1, Matches: matches})
	}

	reverseOffset := len(schedule)
	for _, w := range schedule[:rounds] {
		reversed := domain.Week{WeekNumber: w.WeekNumber + reverseOffset}
		for idx, m := range w.Matches {
			reversed.Matches = append(reversed.Matches, domain.Match{
				ID:         matchID(reversed.WeekNumber, idx+1),
				HomeTeamID: m.AwayTeamID,
				AwayTeamID: m.HomeTeamID,
			})
		}
		schedule = append(schedule, reversed)
	}

	return schedule
}
