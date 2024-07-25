package handlers_http_private_health_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

func (h *Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, nil))
}
