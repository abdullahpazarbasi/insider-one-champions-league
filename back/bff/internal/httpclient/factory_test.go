package httpclient

import (
	"testing"
	"time"
)

func TestNewDefaultTimeout(t *testing.T) {
	client := New(0)
	if client.Timeout != 2*time.Second {
		t.Fatalf("expected default timeout, got %v", client.Timeout)
	}

	custom := New(5 * time.Second)
	if custom.Timeout != 5*time.Second {
		t.Fatalf("expected custom timeout, got %v", custom.Timeout)
	}
}
