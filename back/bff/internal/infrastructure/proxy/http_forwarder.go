package proxy

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"
	"github.com/labstack/echo/v4"
)

func NewHTTPForwarder(client ports.HTTPClient) (*HTTPForwarder, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is required")
	}

	return &HTTPForwarder{client: client}, nil
}

func (f *HTTPForwarder) ForwardGet(ctx context.Context, target string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	copyRequestHeaders(req.Header, headers)
	return f.client.Do(req)
}

func (f *HTTPForwarder) ForwardPost(ctx context.Context, target string, body []byte, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	copyRequestHeaders(req.Header, headers)
	if req.Header.Get(echo.HeaderContentType) == "" {
		req.Header.Set(echo.HeaderContentType, "application/json")
	}

	return f.client.Do(req)
}

func copyRequestHeaders(destination http.Header, source http.Header) {
	for key, values := range source {
		for _, value := range values {
			destination.Add(key, value)
		}
	}
}
