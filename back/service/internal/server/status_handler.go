package server

import (
	"net/http"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
	"github.com/labstack/echo/v4"
)

func (h SimulationHandler) Status(c echo.Context) error {
	return c.JSON(http.StatusOK, domain.StatusResponse{Service: "service", Status: "ok"})
}
