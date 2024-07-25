package pkg_http

import (
	"context"
	"net/http"

	"github.com/teyz/go-svc-template/pkg/errors"
)

func TranslateError(ctx context.Context, err error) (int, interface{}) {
	switch {
	case errors.IsNotFoundError(err):
		return http.StatusNotFound, NewHTTPResponse(http.StatusNotFound, err.Error(), nil)
	case errors.IsResourceAlreadyCreatedError(err):
		return http.StatusConflict, NewHTTPResponse(http.StatusConflict, err.Error(), nil)
	case errors.IsBadRequestError(err):
		return http.StatusBadRequest, NewHTTPResponse(http.StatusBadRequest, err.Error(), nil)
	case errors.IsUnauthorizedError(err):
		return http.StatusUnauthorized, NewHTTPResponse(http.StatusUnauthorized, err.Error(), nil)
	default:
		return http.StatusInternalServerError, NewHTTPResponse(http.StatusInternalServerError, MessageInternalServerError, nil)
	}
}
