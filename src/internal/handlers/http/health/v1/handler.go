package handlers_http_private_health_v1

import (
	"context"
)

type Handler struct{}

func NewHandler(_ context.Context) *Handler {
	return &Handler{}
}
