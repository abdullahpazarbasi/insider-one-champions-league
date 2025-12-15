package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewStatusHandler(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		response := service.Status(c.Request().Context())
		return c.JSON(http.StatusOK, response)
	}
}
