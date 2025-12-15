package simulation

func NewDefaultSimulationService() Service {
	provider := NewStaticTeamProvider()
	validator := StrictTeamValidator{}
	idAssigner := SequentialTeamIDAssigner{}
	scheduleGenerator := RoundRobinScheduleGenerator{}
	weekBuilder := NewWeekBuilderService(scheduleGenerator)
	scoreCalculator := HashScoreCalculator{}
	weekSimulator := NewDeterministicWeekSimulator(scoreCalculator)
	standingsCalculator := TableStandingsCalculator{}
	weekChecker := MatchScoreWeekCompletionChecker{}
	chanceDecider := NewWeekFourChampionChanceDecider(weekChecker)
	rounder := TwoDecimalRounder{}
	chanceCalculator := NewWeightedChampionChanceCalculator(rounder)

	return NewSeasonSimulationService(
		provider,
		validator,
		idAssigner,
		weekBuilder,
		TargetWeekResolverService{},
		weekSimulator,
		standingsCalculator,
		chanceDecider,
		chanceCalculator,
	)
}
