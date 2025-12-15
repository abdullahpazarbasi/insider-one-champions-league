package proxy

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHTTPForwarderValidation(t *testing.T) {
	if _, err := NewHTTPForwarder(nil); err == nil {
		t.Fatalf("expected error when client is missing")
	}
}

func TestForwardGet(t *testing.T) {
	receivedHeaders := http.Header{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedHeaders = r.Header
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	forwarder, err := NewHTTPForwarder(server.Client())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	headers := http.Header{"Authorization": []string{"Bearer token"}}
	resp, err := forwarder.ForwardGet(context.Background(), server.URL, headers)
	if err != nil {
		t.Fatalf("expected request to succeed: %v", err)
	}
	resp.Body.Close()

	if receivedHeaders.Get("Authorization") != "Bearer token" {
		t.Fatalf("header was not forwarded")
	}
}

func TestForwardPostSetsContentType(t *testing.T) {
	var body string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		body = string(data)
		w.WriteHeader(http.StatusCreated)
	}))
	t.Cleanup(server.Close)

	forwarder, err := NewHTTPForwarder(server.Client())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	resp, err := forwarder.ForwardPost(context.Background(), server.URL, []byte(`{"ok":true}`), http.Header{})
	if err != nil {
		t.Fatalf("expected request to succeed: %v", err)
	}
	resp.Body.Close()

	if body != `{"ok":true}` {
		t.Fatalf("unexpected body forwarded: %s", body)
	}

	if resp.Request.Header.Get("Content-Type") == "" {
		t.Fatalf("expected content type to be set")
	}
}

func TestForwardGetInvalidURL(t *testing.T) {
	forwarder, _ := NewHTTPForwarder(&http.Client{})
	if _, err := forwarder.ForwardGet(context.Background(), ":://bad", nil); err == nil {
		t.Fatalf("expected error for invalid url")
	}
}

func TestForwardPostInvalidURL(t *testing.T) {
	forwarder, _ := NewHTTPForwarder(&http.Client{})
	if _, err := forwarder.ForwardPost(context.Background(), ":://bad", []byte("{}"), http.Header{}); err == nil {
		t.Fatalf("expected error for invalid url")
	}
}
