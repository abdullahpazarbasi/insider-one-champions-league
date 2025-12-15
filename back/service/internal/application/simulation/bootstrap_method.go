package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (s SeasonSimulationService) Bootstrap() domain.SimulationResponse {
	defaultTeams := s.teamProvider.DefaultTeams()
	weeks := s.weekBuilder.EnsureWeeks(nil, defaultTeams)
	simulatedWeeks := s.weekSimulator.ApplySimulation(weeks, defaultTeams, 0)
	standings := s.standingsCalculator.ComputeStandings(simulatedWeeks, defaultTeams)
	championChances := s.resolveChampionChances(simulatedWeeks, standings, defaultTeams)

	return domain.SimulationResponse{
		Teams:           defaultTeams,
		Weeks:           simulatedWeeks,
		Standings:       standings,
		ChampionChances: championChances,
	}
}
