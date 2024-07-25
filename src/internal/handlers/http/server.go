package handlers_http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/teyz/go-svc-template/internal/handlers"
	handlers_http_private_health_v1 "github.com/teyz/go-svc-template/internal/handlers/http/health/v1"
	handlers_http_private_example_v1 "github.com/teyz/go-svc-template/internal/handlers/http/private/example/v1"
	service_v1 "github.com/teyz/go-svc-template/internal/service/v1"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

type httpServer struct {
	router  *echo.Echo
	config  pkg_http.HTTPServerConfig
	service service_v1.ExampleStoreService
}

func NewServer(ctx context.Context, cfg pkg_http.HTTPServerConfig, service service_v1.ExampleStoreService) (handlers.Server, error) {
	return &httpServer{
		router:  echo.New(),
		config:  cfg,
		service: service,
	}, nil
}

func (s *httpServer) Setup(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Setup: Setting up HTTP server...")

	// setup handlers
	privateHealhV1Handlers := handlers_http_private_health_v1.NewHandler(ctx, s.service)
	privateExampleV1Handlers := handlers_http_private_example_v1.NewHandler(ctx, s.service)

	// setup middlewares
	//s.router.Use(middleware.Logger())
	s.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.CORS())

	// health endpoints
	s.router.GET("/health", privateHealhV1Handlers.HealthCheck)

	// private endpoints
	privateV1 := s.router.Group("/private/v1")

	// example endpoints
	examplesV1 := privateV1.Group("/examples")
	examplesV1.GET("/", privateExampleV1Handlers.GetExamples)
	examplesV1.POST("/", privateExampleV1Handlers.CreateExample)
	examplesV1.GET("/:id", privateExampleV1Handlers.GetExampleByID)

	return nil
}

func (s *httpServer) Start(ctx context.Context) error {
	log.Info().
		Uint16("port", s.config.Port).
		Msg("handlers.http.httpServer.Start: Starting HTTP server...")

	return s.router.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *httpServer) Stop(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Stop: Stopping HTTP server...")

	return s.router.Shutdown(ctx)
}
