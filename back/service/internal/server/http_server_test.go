package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/application/simulation"
	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
)

type stubSimulationService struct {
	response domain.SimulationResponse
	err      error
}

func (s stubSimulationService) Bootstrap() domain.SimulationResponse {
	return s.response
}

func (s stubSimulationService) Simulate(domain.SimulationRequest) (domain.SimulationResponse, error) {
	return s.response, s.err
}

func TestStatusEndpoint(t *testing.T) {
	handler := NewSimulationHandler(simulation.NewDefaultSimulationService())
	srv := NewHTTPServer(handler)
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}
}

func TestBootstrapEndpointReturnsPayload(t *testing.T) {
	resp := domain.SimulationResponse{Teams: []domain.Team{{ID: "t"}}}
	handler := NewSimulationHandler(stubSimulationService{response: resp})
	srv := NewHTTPServer(handler)
	req := httptest.NewRequest(http.MethodGet, "/bootstrap", nil)
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}

	var body domain.SimulationResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if len(body.Teams) != 1 {
		t.Fatalf("expected stub response to flow through handler")
	}
}

func TestSimulateEndpointHandlesErrors(t *testing.T) {
	handler := NewSimulationHandler(stubSimulationService{err: fmt.Errorf("boom")})
	srv := NewHTTPServer(handler)
	req := httptest.NewRequest(http.MethodPost, "/simulate", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected bad request when service fails, got %d", rec.Code)
	}
}

func TestSimulateEndpointValidatesRequestBody(t *testing.T) {
	handler := NewSimulationHandler(stubSimulationService{})
	srv := NewHTTPServer(handler)
	req := httptest.NewRequest(http.MethodPost, "/simulate", bytes.NewReader([]byte("not-json")))
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected bad request for invalid json, got %d", rec.Code)
	}
}

func TestSimulateEndpointReturnsSuccess(t *testing.T) {
	resp := domain.SimulationResponse{Weeks: []domain.Week{{Matches: []domain.Match{{HomeScore: intPtr(1), AwayScore: intPtr(1)}}}}}
	handler := NewSimulationHandler(stubSimulationService{response: resp})
	srv := NewHTTPServer(handler)

	payload, _ := json.Marshal(domain.SimulationRequest{})
	req := httptest.NewRequest(http.MethodPost, "/simulate", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	srv.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}

	var body domain.SimulationResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	if len(body.Weeks) == 0 || body.Weeks[0].Matches[0].HomeScore == nil {
		t.Fatalf("expected stubbed weeks to be returned")
	}
}

func TestHTTPServerStartLifecycle(t *testing.T) {
	handler := NewSimulationHandler(stubSimulationService{})
	srv := NewHTTPServer(handler)

	done := make(chan struct{})
	go func() {
		_ = srv.Start(":0")
		close(done)
	}()

	srv.engine.HideBanner = true
	_ = srv.engine.Close()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatalf("server did not shut down in time")
	}
}

func intPtr(value int) *int {
	return &value
}
