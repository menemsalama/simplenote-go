// NOTE: this code is partially copied from chi examples code
// TODO: move ErrorResponse and revise it

package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse generic api-error type
type ErrorResponse struct {
	Errors  []string `json:"errors,omitempty"`
	Code    int      `json:"code,omitempty"`
	Status  string   `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
}

// Append adds errors as strings
func (e *ErrorResponse) Append(errs ...error) {
	strErrs := make([]string, len(errs))

	for i, err := range errs {
		strErrs[i] = err.Error()
		log.Println("Error: ", err)
	}

	e.Errors = strErrs
}

// Render status code writer for the error
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}

// InvalidRequestError return 400 responses error
func InvalidRequestError(errs ...error) render.Renderer {
	apiErr := &ErrorResponse{
		Code:   400,
		Status: "Invalid request",
	}

	apiErr.Append(errs...)

	return apiErr
}

// UnauthorizedError return 401 responses error
func UnauthorizedError(msg string, errs ...error) render.Renderer {
	apiErr := &ErrorResponse{
		Code:    401,
		Status:  "Invalid authentication",
		Message: msg,
	}

	apiErr.Append(errs...)

	return apiErr
}
