package config

import "github.com/joho/godotenv"

func LoadEnvironments(paths ...string) {
	for _, path := range paths {
		_ = godotenv.Load(path)
	}
}
