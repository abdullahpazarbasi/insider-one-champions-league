package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/config"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/httpclient"
	infraProxy "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/infrastructure/proxy"
	infraUpstream "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/infrastructure/upstream"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/server"
)

func main() {
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load()

	cfg, err := config.Load(config.NewOSEnvProvider())
	if err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	client := httpclient.New(0)

	forwarder, err := infraProxy.NewHTTPForwarder(client)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	checker, err := infraUpstream.NewHTTPUpstreamChecker(client, cfg.ServiceBaseURL+"/status")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	srv, err := server.New(cfg, forwarder, checker)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := srv.Start(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start bff: %v", err)
	}
}
