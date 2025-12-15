package upstream

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/domain/status"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"
)

func NewHTTPUpstreamChecker(client ports.HTTPClient, target string) (*HTTPUpstreamChecker, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is required")
	}

	parsed, err := url.ParseRequestURI(target)
	if err != nil {
		return nil, fmt.Errorf("invalid upstream url: %w", err)
	}

	return &HTTPUpstreamChecker{client: client, statusURL: parsed.String()}, nil
}

func (c *HTTPUpstreamChecker) Check(ctx context.Context) (*status.UpstreamStatus, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.statusURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("upstream returned non-2xx status")
	}

	return &status.UpstreamStatus{Service: "service", Status: "ok"}, nil
}
