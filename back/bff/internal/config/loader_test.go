package config

import "testing"

type stubProvider struct {
	values map[string]string
}

func (s stubProvider) Get(key string) string {
	return s.values[key]
}

func TestLoadConfiguration(t *testing.T) {
	provider := stubProvider{values: map[string]string{
		"PORT":             "8081",
		"SERVICE_BASE_URL": "http://example.com/api/",
	}}

	cfg, err := Load(provider)
	if err != nil {
		t.Fatalf("expected configuration to load: %v", err)
	}

	if cfg.Port != "8081" {
		t.Fatalf("unexpected port: %s", cfg.Port)
	}

	if cfg.ServiceBaseURL != "http://example.com/api" {
		t.Fatalf("unexpected base url: %s", cfg.ServiceBaseURL)
	}
}

func TestLoadConfigurationValidation(t *testing.T) {
	_, err := Load(nil)
	if err == nil {
		t.Fatalf("expected error for nil provider")
	}

	_, err = Load(stubProvider{values: map[string]string{"SERVICE_BASE_URL": "http://example.com"}})
	if err == nil {
		t.Fatalf("expected error when port is missing")
	}

	_, err = Load(stubProvider{values: map[string]string{"PORT": "8080"}})
	if err == nil {
		t.Fatalf("expected error when service base url is missing")
	}
}

func TestOSEnvProvider(t *testing.T) {
	t.Setenv("SAMPLE_KEY", "value")
	provider := NewOSEnvProvider()

	if provider.Get("SAMPLE_KEY") != "value" {
		t.Fatalf("expected env value to be returned")
	}
}
