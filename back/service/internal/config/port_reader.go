package config

import (
	"errors"
	"os"
	"strings"
)

func ReadPortFromEnv() (string, error) {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		return "", errors.New("PORT environment variable is required")
	}
	return port, nil
}
