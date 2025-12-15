package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h SimulationHandler) Bootstrap(c echo.Context) error {
	payload := h.service.Bootstrap()
	return c.JSON(http.StatusOK, payload)
}
