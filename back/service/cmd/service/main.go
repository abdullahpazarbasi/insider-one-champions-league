package main

import (
	"log"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/application/simulation"
	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/config"
	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/server"
)

func main() {
	config.LoadEnvironments(".env.local", ".env")
	port, err := config.ReadPortFromEnv()
	if err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	simulationService := simulation.NewDefaultSimulationService()
	handler := server.NewSimulationHandler(simulationService)
	httpServer := server.NewHTTPServer(handler)

	if err := httpServer.Start(":" + port); err != nil {
		log.Fatalf("failed to start service: %v", err)
	}
}
