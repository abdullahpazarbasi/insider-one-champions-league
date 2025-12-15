package simulation

import "fmt"

func matchID(weekNumber, matchNumber int) string {
	return fmt.Sprintf("w%d-m%d", weekNumber, matchNumber)
}
