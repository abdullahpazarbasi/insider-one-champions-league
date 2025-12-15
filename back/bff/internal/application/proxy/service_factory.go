package proxy

import (
	"fmt"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"
)

func NewService(forwarder ports.Forwarder) (Service, error) {
	if forwarder == nil {
		return Service{}, fmt.Errorf("forwarder is required")
	}

	return Service{forwarder: forwarder}, nil
}
