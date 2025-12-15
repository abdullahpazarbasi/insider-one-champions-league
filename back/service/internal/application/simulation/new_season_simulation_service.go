package simulation

func NewSeasonSimulationService(
	teamProvider TeamProvider,
	teamValidator TeamValidator,
	teamIDAssigner TeamIDAssigner,
	weekBuilder WeekBuilder,
	targetWeekResolver TargetWeekResolver,
	weekSimulator WeekSimulator,
	standingsCalculator StandingsCalculator,
	championChanceDecider ChampionChanceDecider,
	championChanceCalculator ChampionChanceCalculator,
) SeasonSimulationService {
	return SeasonSimulationService{
		teamProvider:             teamProvider,
		teamValidator:            teamValidator,
		teamIDAssigner:           teamIDAssigner,
		weekBuilder:              weekBuilder,
		targetWeekResolver:       targetWeekResolver,
		weekSimulator:            weekSimulator,
		standingsCalculator:      standingsCalculator,
		championChanceDecider:    championChanceDecider,
		championChanceCalculator: championChanceCalculator,
	}
}
