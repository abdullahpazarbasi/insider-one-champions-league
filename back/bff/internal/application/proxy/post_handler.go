package proxy

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPostHandler(service Service, target string) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		resp, err := service.Post(c.Request().Context(), target, body, c.Request().Header)
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{"error": "upstream unavailable"})
		}
		defer resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			return c.JSON(resp.StatusCode, map[string]string{"error": "upstream error"})
		}

		c.Response().WriteHeader(resp.StatusCode)
		_, copyErr := io.Copy(c.Response(), resp.Body)
		return copyErr
	}
}
