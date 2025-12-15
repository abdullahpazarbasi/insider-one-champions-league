package config

import "os"

type OSEnvProvider struct{}

func NewOSEnvProvider() OSEnvProvider {
	return OSEnvProvider{}
}

func (OSEnvProvider) Get(key string) string {
	return os.Getenv(key)
}
