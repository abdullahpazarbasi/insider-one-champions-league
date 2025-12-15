package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/config"
	infraProxy "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/infrastructure/proxy"
	infraUpstream "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/infrastructure/upstream"
)

type passthroughClient struct {
	handler http.Handler
}

func (p passthroughClient) Do(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	p.handler.ServeHTTP(rr, req)
	return rr.Result(), nil
}

func TestServerRoutes(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/status":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"service":"service","status":"ok"}`))
		case "/bootstrap":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"teams":[],"weeks":[]}`))
		case "/simulate":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"ok":true}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(upstream.Close)

	cfg := config.Config{Port: "8080", ServiceBaseURL: upstream.URL}
	client := upstream.Client()

	forwarder, err := infraProxy.NewHTTPForwarder(client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	checker, err := infraUpstream.NewHTTPUpstreamChecker(client, upstream.URL+"/status")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	srv, err := New(cfg, forwarder, checker)
	if err != nil {
		t.Fatalf("failed to configure server: %v", err)
	}

	t.Run("status", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/status", nil)
		rec := httptest.NewRecorder()

		srv.router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("unexpected status code: %d", rec.Code)
		}
	})

	t.Run("bootstrap", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/bootstrap", nil)
		rec := httptest.NewRecorder()

		srv.router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("unexpected status code: %d", rec.Code)
		}
	})

	t.Run("simulate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/simulate", nil)
		rec := httptest.NewRecorder()

		srv.router.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("unexpected status code: %d", rec.Code)
		}
	})
}

func TestServerValidation(t *testing.T) {
	cfg := config.Config{Port: "8080", ServiceBaseURL: ":://"}
	_, err := New(cfg, nil, nil)
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestServerRequiresDependencies(t *testing.T) {
	cfg := config.Config{Port: "8080", ServiceBaseURL: "http://example.com"}

	if _, err := New(cfg, nil, nil); err == nil {
		t.Fatalf("expected error when forwarder is nil")
	}

	forwarder, _ := infraProxy.NewHTTPForwarder(&http.Client{})
	if _, err := New(cfg, forwarder, nil); err == nil {
		t.Fatalf("expected error when checker is nil")
	}
}

func TestServerInvalidBaseURL(t *testing.T) {
	client := &http.Client{}
	forwarder, _ := infraProxy.NewHTTPForwarder(client)
	checker, _ := infraUpstream.NewHTTPUpstreamChecker(passthroughClient{handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})}, "http://example.com/status")

	_, err := New(config.Config{Port: "8080", ServiceBaseURL: ":://"}, forwarder, checker)
	if err == nil {
		t.Fatalf("expected error for invalid base url")
	}
}

func TestServerStartAndShutdown(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(upstream.Close)

	cfg := config.Config{Port: "0", ServiceBaseURL: upstream.URL}
	client := upstream.Client()

	forwarder, err := infraProxy.NewHTTPForwarder(client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	checker, err := infraUpstream.NewHTTPUpstreamChecker(client, upstream.URL+"/status")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	srv, err := New(cfg, forwarder, checker)
	if err != nil {
		t.Fatalf("failed to configure server: %v", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- srv.Start(":0")
	}()

	time.Sleep(50 * time.Millisecond)
	shutdownErr := srv.router.Shutdown(context.Background())
	if shutdownErr != nil {
		t.Fatalf("failed to shutdown server: %v", shutdownErr)
	}

	if err := <-done; err != nil && err != http.ErrServerClosed {
		t.Fatalf("unexpected start error: %v", err)
	}
}
