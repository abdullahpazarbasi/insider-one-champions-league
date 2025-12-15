package upstream

import "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"

type HTTPUpstreamChecker struct {
	client    ports.HTTPClient
	statusURL string
}
