package handlers_http_private_health_v1

import (
	"context"

	service_v1 "github.com/teyz/go-svc-template/internal/service/v1"
)

type Handler struct {
	service service_v1.ExampleStoreService
}

func NewHandler(_ context.Context, service service_v1.ExampleStoreService) *Handler {
	return &Handler{
		service: service,
	}
}
