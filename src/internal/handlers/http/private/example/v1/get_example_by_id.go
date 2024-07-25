package handlers_http_private_example_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

type GetExampleByIDResponse struct {
	Example *entities_example_v1.Example `json:"example"`
}

func (h *Handler) GetExampleByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		log.Error().Msg("handlers.http.private.example.v1.get_example.Handler.GetExampleByID: can not get id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	example, err := h.service.GetExampleByID(ctx, id)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetExampleByIDResponse{
		Example: &entities_example_v1.Example{
			ID:          example.ID,
			Description: example.Description,
			CreatedAt:   example.CreatedAt,
			UpdatedAt:   example.UpdatedAt,
		},
	}))
}
