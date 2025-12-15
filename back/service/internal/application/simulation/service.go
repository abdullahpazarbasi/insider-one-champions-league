package simulation

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"

type Service interface {
	Bootstrap() domain.SimulationResponse
	Simulate(request domain.SimulationRequest) (domain.SimulationResponse, error)
}
