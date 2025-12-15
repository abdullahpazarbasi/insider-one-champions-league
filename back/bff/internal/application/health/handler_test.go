package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/domain/status"
	"github.com/labstack/echo/v4"
)

func TestStatusHandler(t *testing.T) {
	svc, err := NewService(stubChecker{upstream: &status.UpstreamStatus{Service: "service", Status: "ok"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	if err := NewStatusHandler(svc)(ctx); err != nil {
		t.Fatalf("handler error: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}
}
