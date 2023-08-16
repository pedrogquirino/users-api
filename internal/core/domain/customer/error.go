package customer

import (
	"fmt"
	"net/http"
)

type Error struct {
	Message        string `json:"message"`
	HttpStatusCode int    `json:"http_status_code"`
}

func (e Error) Error() string {
	errorMessage := fmt.Sprintf("error: %s", e.Message)
	return errorMessage
}

var InvalidRequestError = Error{
	Message:        "Invalid request",
	HttpStatusCode: http.StatusBadRequest,
}

var InvalidRequestValidationError = Error{
	Message:        "Invalid request struct",
	HttpStatusCode: http.StatusBadRequest,
}

var InternalError = Error{
	Message:        "Internal error",
	HttpStatusCode: http.StatusInternalServerError,
}
