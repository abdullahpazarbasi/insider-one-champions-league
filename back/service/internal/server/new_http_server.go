package server

import (
	"github.com/labstack/echo/v4"
)

func NewHTTPServer(handler SimulationHandler) HTTPServer {
	router := echo.New()
	router.HideBanner = true

	router.GET("/status", handler.Status)
	router.GET("/bootstrap", handler.Bootstrap)
	router.POST("/simulate", handler.Simulate)

	return HTTPServer{engine: router}
}
