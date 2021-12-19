package httputils

import (
	"net/http"

	"github.com/go-chi/render"
)

const (
	titleErrorRenderingResponse = "error rendering response"
	titleResourceNotFound       = "error resource not found"
	titleRequestBodyMissing     = "error request body is missing"
	titleControlNotImplemented  = "error control interface not implemented"
)

// ErrResponse structure response with message.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Status    string `json:"status"` // user-level status message
	Code      int64  `json:"code"`   // application-specific error code
	ErrorText string `json:"error"`  // application-level error message, for debugging
}

// Render renders only status code.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// Error returns error text message.
func (e *ErrResponse) Error() string {
	return e.ErrorText
}

// ErrBadRequest create 400 response.
func ErrBadRequest(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Status:         http.StatusText(http.StatusBadRequest),
		ErrorText:      err.Error(),
	}
}

// ErrNotFound create 404 response.
func ErrNotFound(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		Status:         http.StatusText(http.StatusNotFound),
		ErrorText:      err.Error(),
	}
}

// ErrRender create 422 response.
func ErrRender(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Status:         titleErrorRenderingResponse,
		ErrorText:      err.Error(),
	}
}

// ErrInternalServer create 500 response.
func ErrInternalServer(err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		Status:         http.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

// ErrNotFoundDefault default message with 404.
var ErrNotFoundDefault = &ErrResponse{HTTPStatusCode: http.StatusNotFound, Status: http.StatusText(http.StatusNotFound), ErrorText: titleResourceNotFound}

// ErrMissingBody default message with 400.
var ErrMissingBody = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, Status: http.StatusText(http.StatusBadRequest), ErrorText: titleRequestBodyMissing}

// ErrCtlNotImplemented default message with 500.
var ErrCtlNotImplemented = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, Status: http.StatusText(http.StatusInternalServerError), ErrorText: titleControlNotImplemented}
