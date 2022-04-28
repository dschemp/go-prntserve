package response

import (
    "github.com/go-chi/render"
    "net/http"
)

// ErrNotImplementedYet returns a render.Renderer prefilled with a generic not implemented yet message.
// This should be used, if the specific action has yet to be implemented.
func ErrNotImplementedYet() render.Renderer {
	return &ServerResponse{
		HTTPStatusCode: http.StatusNotImplemented,
		Message:        "This has not been implemented yet!",
		Code:           "NOT_IMPLEMENTED",
	}
}

// ErrInternalServerError returns a render.Renderer prefilled with a generic internal server error message.
// This should be used, if the error is known and can be passed to the user.
//
// The passed error is shown to the user as an additional text.
func ErrInternalServerError(err error) render.Renderer {
	return &ServerResponse{
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
	return &ServerResponse{
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
	return &ServerResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		Message:        "You are not authorized to access this resource!",
		Code:           "UNAUTHORIZED",
		ErrorText:      err.Error(),
	}
}

// ErrForbidden returns a render.Renderer prefilled with a generic forbidden message.
// This should be used, if the identity of the user is known, otherwise use ErrUnauthorized.
//
// The passed error is shown to the user as an additional text.
func ErrForbidden(err error) render.Renderer {
	return &ServerResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		Message:        "You are not allowed to access this resource!",
		Code:           "FORBIDDEN",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound returns a render.Renderer prefilled with a generic not found message.
// This should be used, if a resource could not be found, for example if it does not exist.
func ErrNotFound() render.Renderer {
	return &ServerResponse{
		HTTPStatusCode: http.StatusNotFound,
		Message:        "The requested resource could not be found!",
		Code:           "NOT_FOUND",
	}
}

// ErrMethodNotAllowed returns a render.Renderer prefilled with a generic method not allowed message.
// This should be used, if a resource cannot be accessed with this certain method
func ErrMethodNotAllowed() render.Renderer {
	return &ServerResponse{
		HTTPStatusCode: http.StatusMethodNotAllowed,
		Message:        "You are trying to access this resource with an unsupported method!",
		Code:           "METHOD_NOT_ALLOWED",
	}
}

// ErrBadRequest returns a render.Renderer prefilled with a generic bad request message.
// This should be used, if the client provided invalid request data.
func ErrBadRequest(err error) render.Renderer {
	return &ServerResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Message:        "The request you try to perform is invalid!",
		Code:           "BAD_REQUEST",
		ErrorText:      err.Error(),
	}
}

// ErrConflict returns a render.Renderer prefilled with a generic conflict message.
// This should be used, if there is some sort of user initiated conflict.
func ErrConflict(err error) render.Renderer {
	return &ServerResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		Message:        "The request conflicts with the server!",
		Code:           "CONFLICT",
		ErrorText:      err.Error(),
	}
}
