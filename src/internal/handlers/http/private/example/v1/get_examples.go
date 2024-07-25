package handlers_http_private_example_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
	pkg_http "github.com/teyz/go-svc-template/pkg/http"
)

type GetExamplesResponse struct {
	Examples []*entities_example_v1.Example `json:"examples"`
}

func (h *Handler) GetExamples(c echo.Context) error {
	ctx := c.Request().Context()

	examples, err := h.service.GetExamples(ctx)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	examplesResp := make([]*entities_example_v1.Example, 0, len(examples))

	for _, example := range examplesResp {
		examplesResp = append(examplesResp, &entities_example_v1.Example{
			ID:          example.ID,
			Description: example.Description,
			CreatedAt:   example.CreatedAt,
			UpdatedAt:   example.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetExamplesResponse{
		Examples: examplesResp,
	}))
}
