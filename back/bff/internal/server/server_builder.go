package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/application/health"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/application/proxy"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/config"
	"github.com/abdullahpazarbasi/insider-one-champions-league/bff/internal/ports"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(cfg config.Config, forwarder ports.Forwarder, checker health.UpstreamChecker) (*Server, error) {
	if forwarder == nil {
		return nil, fmt.Errorf("forwarder is required")
	}
	if checker == nil {
		return nil, fmt.Errorf("upstream checker is required")
	}

	if _, err := url.ParseRequestURI(cfg.ServiceBaseURL); err != nil {
		return nil, fmt.Errorf("invalid service base url: %w", err)
	}

	router := echo.New()
	router.HideBanner = true
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"https://iocl.local",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

	router.Use(middleware.Secure())

	proxyService, err := proxy.NewService(forwarder)
	if err != nil {
		return nil, err
	}

	healthService, err := health.NewService(checker)
	if err != nil {
		return nil, err
	}

	router.GET("/status", health.NewStatusHandler(healthService))
	router.GET("/api/bootstrap", proxy.NewGetHandler(proxyService, cfg.ServiceBaseURL+"/bootstrap"))
	router.POST("/api/simulate", proxy.NewPostHandler(proxyService, cfg.ServiceBaseURL+"/simulate"))

	return &Server{router: router}, nil
}
