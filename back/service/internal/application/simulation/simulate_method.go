package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

func (s SeasonSimulationService) Simulate(request domain.SimulationRequest) (domain.SimulationResponse, error) {
	if err := s.teamValidator.ValidateTeams(request.Teams); err != nil {
		return domain.SimulationResponse{}, err
	}

	teams := s.teamIDAssigner.EnsureTeamIDs(request.Teams)
	weeks := s.weekBuilder.EnsureWeeks(request.Weeks, teams)

	targetWeekIndex, err := s.targetWeekResolver.ResolveTargetWeekIndex(request.TargetWeekIndex, weeks)
	if err != nil {
		return domain.SimulationResponse{}, err
	}

	simulatedWeeks := s.weekSimulator.ApplySimulation(weeks, teams, targetWeekIndex)
	standings := s.standingsCalculator.ComputeStandings(simulatedWeeks, teams)
	championChances := s.resolveChampionChances(simulatedWeeks, standings, teams)

	return domain.SimulationResponse{
		Teams:           teams,
		Weeks:           simulatedWeeks,
		Standings:       standings,
		ChampionChances: championChances,
	}, nil
}
