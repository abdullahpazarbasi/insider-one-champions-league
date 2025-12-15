package proxy

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewGetHandler(service Service, target string) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := service.Get(c.Request().Context(), target, c.Request().Header)
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{"error": "upstream unavailable"})
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			return c.JSON(resp.StatusCode, map[string]string{"error": "upstream error"})
		}

		copyHeaders(resp.Header, c)
		c.Response().WriteHeader(resp.StatusCode)
		_, copyErr := io.Copy(c.Response(), resp.Body)
		return copyErr
	}
}
