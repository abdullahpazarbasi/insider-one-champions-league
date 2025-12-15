package health

import (
	"context"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/domain/status"
)

type UpstreamChecker interface {
	Check(ctx context.Context) (*status.UpstreamStatus, error)
}
