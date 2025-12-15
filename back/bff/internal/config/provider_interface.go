package config

type Provider interface {
	Get(key string) string
}
