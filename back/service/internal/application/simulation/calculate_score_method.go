package simulation

import (
	"fmt"
	"hash/fnv"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

func (HashScoreCalculator) CalculateScore(home domain.Team, away domain.Team, weekNumber int, matchIndex int) (int, int) {
	seed := fmt.Sprintf("%s-%s-%d-%d-%d-%d", home.ID, away.ID, home.Strength, away.Strength, weekNumber, matchIndex)
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(seed))
	value := hash.Sum32()

	bias := home.Strength - away.Strength

	homeBase := int(value % 3)
	awayBase := int((value / 7) % 3)

	if bias > 20 {
		homeBase++
	}
	if bias < -20 {
		awayBase++
	}

	if bias > 40 {
		homeBase++
	}
	if bias < -40 {
		awayBase++
	}

	return clampScore(homeBase), clampScore(awayBase)
}
