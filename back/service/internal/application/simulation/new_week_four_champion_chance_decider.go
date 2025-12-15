package simulation

func NewWeekFourChampionChanceDecider(checker WeekCompletionChecker) WeekFourChampionChanceDecider {
	return WeekFourChampionChanceDecider{weekChecker: checker}
}
