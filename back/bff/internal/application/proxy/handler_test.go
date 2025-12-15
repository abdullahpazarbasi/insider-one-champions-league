package proxy

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

type stubForwarder struct {
	response *http.Response
	err      error
}

func (s stubForwarder) ForwardGet(_ context.Context, _ string, _ http.Header) (*http.Response, error) {
	return s.response, s.err
}

func (s stubForwarder) ForwardPost(_ context.Context, _ string, _ []byte, _ http.Header) (*http.Response, error) {
	return s.response, s.err
}

func TestGetHandlerSuccess(t *testing.T) {
	handlerResponse := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
	}

	service, err := NewService(stubForwarder{response: handlerResponse})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/bootstrap", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	if err := NewGetHandler(service, "http://example.com")(ctx); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}

	if rec.Header().Get("Content-Type") == "" {
		t.Fatalf("expected headers to be forwarded")
	}
}

func TestGetHandlerErrors(t *testing.T) {
	service, _ := NewService(stubForwarder{err: errors.New("boom")})
	req := httptest.NewRequest(http.MethodGet, "/api/bootstrap", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	_ = NewGetHandler(service, "http://example.com")(ctx)
	if rec.Code != http.StatusBadGateway {
		t.Fatalf("expected bad gateway for upstream errors, got %d", rec.Code)
	}

	upstreamResp := &http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(strings.NewReader(""))}
	service, _ = NewService(stubForwarder{response: upstreamResp})
	rec = httptest.NewRecorder()
	ctx = echo.New().NewContext(req, rec)

	_ = NewGetHandler(service, "http://example.com")(ctx)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected upstream status, got %d", rec.Code)
	}
}

func TestPostHandler(t *testing.T) {
	handlerResponse := &http.Response{StatusCode: http.StatusCreated, Body: io.NopCloser(strings.NewReader("created"))}
	service, _ := NewService(stubForwarder{response: handlerResponse})

	req := httptest.NewRequest(http.MethodPost, "/api/simulate", strings.NewReader("{}"))
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	if err := NewPostHandler(service, "http://example.com")(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status: %d", rec.Code)
	}
}

func TestPostHandlerInvalidBody(t *testing.T) {
	service, _ := NewService(stubForwarder{err: errors.New("unavailable")})
	req := httptest.NewRequest(http.MethodPost, "/api/simulate", errReader{})
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	_ = NewPostHandler(service, "http://example.com")(ctx)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected bad request for unreadable body, got %d", rec.Code)
	}
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("read error")
}

func (errReader) Close() error { return nil }

func TestPostHandlerUpstreamError(t *testing.T) {
	service, _ := NewService(stubForwarder{err: errors.New("upstream down")})

	req := httptest.NewRequest(http.MethodPost, "/api/simulate", strings.NewReader("{}"))
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	_ = NewPostHandler(service, "http://example.com")(ctx)
	if rec.Code != http.StatusBadGateway {
		t.Fatalf("expected bad gateway when upstream fails, got %d", rec.Code)
	}
}
