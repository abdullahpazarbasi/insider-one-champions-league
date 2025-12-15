package proxy

import "testing"

func TestNewServiceValidation(t *testing.T) {
	if _, err := NewService(nil); err == nil {
		t.Fatalf("expected error when forwarder is nil")
	}
}
