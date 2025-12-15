package health

import (
	"context"
	"errors"
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/domain/status"
)

type stubChecker struct {
	upstream *status.UpstreamStatus
	err      error
}

func (s stubChecker) Check(context.Context) (*status.UpstreamStatus, error) {
	return s.upstream, s.err
}

func TestServiceStatusWithUpstream(t *testing.T) {
	upstream := &status.UpstreamStatus{Service: "service", Status: "ok"}
	svc, err := NewService(stubChecker{upstream: upstream})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resp := svc.Status(context.Background())
	if resp.Upstream == nil || resp.Upstream.Status != "ok" {
		t.Fatalf("unexpected upstream status: %+v", resp.Upstream)
	}
}

func TestServiceStatusWithoutUpstream(t *testing.T) {
	svc, err := NewService(stubChecker{err: errors.New("boom")})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resp := svc.Status(context.Background())
	if resp.Upstream != nil {
		t.Fatalf("expected nil upstream when probe fails")
	}
}

func TestNewServiceValidation(t *testing.T) {
	if _, err := NewService(nil); err == nil {
		t.Fatalf("expected validation error")
	}
}
