package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnvironmentsAppliesVariables(t *testing.T) {
	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, "test.env")
	content := []byte("PORT=8080\nCUSTOM=ok")
	if err := os.WriteFile(envFile, content, 0o644); err != nil {
		t.Fatalf("failed to write env file: %v", err)
	}

	os.Unsetenv("CUSTOM")
	LoadEnvironments(envFile)

	if got := os.Getenv("CUSTOM"); got != "ok" {
		t.Fatalf("expected environment to be loaded, got %s", got)
	}
}

func TestReadPortFromEnvReturnsTrimmedValue(t *testing.T) {
	t.Setenv("PORT", " 9090 ")

	port, err := ReadPortFromEnv()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if port != "9090" {
		t.Fatalf("expected trimmed port, got %s", port)
	}
}

func TestReadPortFromEnvFailsWhenMissing(t *testing.T) {
	os.Unsetenv("PORT")

	_, err := ReadPortFromEnv()
	if err == nil {
		t.Fatalf("expected error when port is missing")
	}
}
