package upstream

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type failingClient struct{}

func (f failingClient) Do(*http.Request) (*http.Response, error) {
	return nil, errors.New("network error")
}

type passthroughClient struct {
	handler http.Handler
}

func (p passthroughClient) Do(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	p.handler.ServeHTTP(rr, req)
	return rr.Result(), nil
}

func TestNewHTTPUpstreamCheckerValidation(t *testing.T) {
	if _, err := NewHTTPUpstreamChecker(nil, "http://example.com"); err == nil {
		t.Fatalf("expected error when client is missing")
	}

	if _, err := NewHTTPUpstreamChecker(passthroughClient{}, ":://"); err == nil {
		t.Fatalf("expected error for invalid url")
	}
}

func TestHTTPUpstreamCheckerSuccess(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	checker, err := NewHTTPUpstreamChecker(passthroughClient{handler: handler}, "http://example.com/status")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	upstream, err := checker.Check(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if upstream == nil || upstream.Status != "ok" {
		t.Fatalf("unexpected upstream response: %+v", upstream)
	}
}

func TestHTTPUpstreamCheckerFailure(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})
	checker, err := NewHTTPUpstreamChecker(passthroughClient{handler: handler}, "http://example.com/status")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := checker.Check(context.Background()); err == nil {
		t.Fatalf("expected error when upstream is unhealthy")
	}

	failing, _ := NewHTTPUpstreamChecker(failingClient{}, "http://example.com/status")
	if _, err := failing.Check(context.Background()); err == nil {
		t.Fatalf("expected network error")
	}
}
