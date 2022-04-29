package response

import (
	"github.com/dschemp/go-prntserve/internal/logging"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ServerResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message"`
	ErrorText      string `json:"details,omitempty"`
}

func (e *ServerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	log.Debug().
		Err(e.Err).
		Int(logging.HTTPStatusCodeFieldName, e.HTTPStatusCode).
		Str(logging.SystemCodeFieldName, e.Code).
		Msg(e.Message)
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// RespondRaw enables writing of binary data. This responds always with status 200 OK.
func RespondRaw(w http.ResponseWriter, r *http.Request, data []byte) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write(data)
}

// FileUploadedSuccessfully returns a render.Renderer intended to be used to signify, that a file has been uploaded successfully.
func FileUploadedSuccessfully() render.Renderer {
	return &ServerResponse{
		HTTPStatusCode: http.StatusOK,
		Message:        "File uploaded successfully!",
		Code:           "UPLOAD_FILE_SUCCESS",
	}
}

// FileDeletedSuccessfully returns a render.Renderer intended to be used to signify, that a file has been deleted successfully.
func FileDeletedSuccessfully() render.Renderer {
	return &ServerResponse{
		HTTPStatusCode: http.StatusOK,
		Message:        "File deleted successfully!",
		Code:           "DELETE_FILE_SUCCESS",
	}
}
