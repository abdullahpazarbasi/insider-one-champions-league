package simulation

// WeekFourChampionChanceDecider waits until week four is complete before computing chances.
type WeekFourChampionChanceDecider struct {
	weekChecker WeekCompletionChecker
}
