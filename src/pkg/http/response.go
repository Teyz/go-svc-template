package pkg_http

import "net/http"

const (
	MessageSuccess                 = "SUCCESS"
	MessageBindError               = "MALFORMATED_PARAMETERS_ERROR"
	MessageValidationError         = "INVALID_PARAMETERS_ERROR"
	MessageInternalServerError     = "INTERNAL_SERVER_ERROR"
	MessageUserAlreadyCreatedError = "CONFLICT_ALREADY_EXIST_ERROR"
	MessageBadRequestError         = "BAD_REQUEST_ERROR"
	MessageNotFoundError           = "NOT_FOUND_ERROR"
	MessageUnauthorizedError       = "UNAUTHORIZED_ERROR"
	MessageForbidenError           = "FORBIDEN_ERROR"
)

type HTTPResponseStatus struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPResponse struct {
	Status HTTPResponseStatus `json:"status"`
	Data   interface{}        `json:"data"`
}

// NewHTTPResponse creates a new HTTPResponse
func NewHTTPResponse(statusCode int, message string, data interface{}) HTTPResponse {
	resp := HTTPResponse{
		Status: HTTPResponseStatus{
			Code:    statusCode,
			Message: message,
		},
		Data: data,
	}

	if statusCode != http.StatusOK && statusCode != http.StatusCreated &&
		statusCode != http.StatusNoContent && statusCode != http.StatusAccepted {
		resp.Status.Error = true
	}

	return resp
}
