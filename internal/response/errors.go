package response

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message"`
	ErrorText      string `json:"details,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrNotImplementedYet returns a render.Renderer prefilled with a generic not implemented yet message.
// This should be used, if the specific action has yet to be implemented.
func ErrNotImplementedYet() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusNotImplemented,
		Message:        "This is not implemented yet!",
		Code:           "NOT_IMPLEMENTED",
	}
}

// ErrInternalServerError returns a render.Renderer prefilled with a generic internal server error message.
// This should be used, if the error is known and can be passed to the user.
//
// The passed error is shown to the user as an additional text.
func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        "The server encountered a problem which it doesn't know how to handle.",
		Code:           "INTERNAL_SERVER_ERROR",
		ErrorText:      err.Error(),
	}
}

// ErrInternalServerErrorWithCustomMessage returns a render.Renderer prefilled with a generic internal server error message.
// This should be used, if the error message is known and can be passed to the user.
func ErrInternalServerErrorWithCustomMessage(message string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        message,
		Code:           "INTERNAL_SERVER_ERROR",
	}
}

// ErrUnauthorized returns a render.Renderer prefilled with a generic unauthorized message.
// This should be used, if the identity of the user is unknown, otherwise use ErrForbidden.
//
// The passed error is shown to the user as an additional text.
func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		Message:        "Unauthorized to access this resource!",
		Code:           "UNAUTHORIZED",
		ErrorText:      err.Error(),
	}
}

// ErrForbidden returns a render.Renderer prefilled with a generic forbidden message.
// This should be used, if the identity of the user is known, otherwise use ErrUnauthorized.
//
// The passed error is shown to the user as an additional text.
func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		Message:        "You are not allowed to access this resource!",
		Code:           "FORBIDDEN",
		ErrorText:      err.Error(),
	}
}
