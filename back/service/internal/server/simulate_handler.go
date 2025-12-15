package server

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahpazarbasi/insider-one-champions-league/service/internal/domain"
	"github.com/labstack/echo/v4"
)

func (h SimulationHandler) Simulate(c echo.Context) error {
	var req domain.SimulationRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	resp, err := h.service.Simulate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
