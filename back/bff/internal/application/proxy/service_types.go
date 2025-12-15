package proxy

import "github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"

type Service struct {
	forwarder ports.Forwarder
}
