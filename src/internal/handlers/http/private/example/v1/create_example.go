package handlers_http_private_example_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

type CreateExampleRequest struct {
	Description string `json:"description"`
}

type CreateExampleResponse struct {
	Example *entities_example_v1.Example `json:"example"`
}

func (h *Handler) CreateExample(c echo.Context) error {
	ctx := c.Request().Context()

	var req CreateExampleRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("handlers.http.private.example.v1.create_example.CreateExample: can not bind request")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	if req.Description == "" {
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	example, err := h.service.CreateExample(ctx, req.Description)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusCreated, pkg_http.NewHTTPResponse(http.StatusCreated, pkg_http.MessageSuccess, CreateExampleResponse{
		Example: &entities_example_v1.Example{
			ID:          example.ID,
			Description: example.Description,
			CreatedAt:   example.CreatedAt,
			UpdatedAt:   example.UpdatedAt,
		},
	}))
}
