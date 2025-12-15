package health

import (
	"context"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/domain/status"
)

func (s Service) Status(ctx context.Context) status.StatusResponse {
	upstream, err := s.checker.Check(ctx)
	if err != nil {
		upstream = nil
	}

	return status.StatusResponse{Service: "bff", Status: "ok", Upstream: upstream}
}
