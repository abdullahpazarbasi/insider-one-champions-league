package simulation

// SeasonSimulationService orchestrates the simulation use cases.
type SeasonSimulationService struct {
	teamProvider             TeamProvider
	teamValidator            TeamValidator
	teamIDAssigner           TeamIDAssigner
	weekBuilder              WeekBuilder
	targetWeekResolver       TargetWeekResolver
	weekSimulator            WeekSimulator
	standingsCalculator      StandingsCalculator
	championChanceDecider    ChampionChanceDecider
	championChanceCalculator ChampionChanceCalculator
}
