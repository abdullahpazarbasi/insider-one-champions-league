package proxy

import "github.com/labstack/echo/v4"

func copyHeaders(from map[string][]string, c echo.Context) {
	for key, values := range from {
		for _, value := range values {
			c.Response().Header().Add(key, value)
		}
	}
}
