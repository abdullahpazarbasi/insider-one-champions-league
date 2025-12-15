package proxy

import "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"

type HTTPForwarder struct {
	client ports.HTTPClient
}
