package server

import "github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/application/simulation"

func NewSimulationHandler(service simulation.Service) SimulationHandler {
	return SimulationHandler{service: service}
}
