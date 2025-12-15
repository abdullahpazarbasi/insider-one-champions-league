package config

import (
	"fmt"
	"strings"
)

func Load(provider Provider) (Config, error) {
	if provider == nil {
		return Config{}, fmt.Errorf("config provider is required")
	}

	port := strings.TrimSpace(provider.Get("PORT"))
	if port == "" {
		return Config{}, fmt.Errorf("PORT environment variable is required")
	}

	upstream := strings.TrimSpace(provider.Get("SERVICE_BASE_URL"))
	if upstream == "" {
		return Config{}, fmt.Errorf("SERVICE_BASE_URL environment variable is required")
	}

	return Config{Port: port, ServiceBaseURL: strings.TrimSuffix(upstream, "/")}, nil
}
