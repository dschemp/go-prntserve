package handler

import (
	"github.com/dschemp/go-prntserve/internal/logging"
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	filePath := chi.URLParam(r, "filepath")

	data, err := GetFileFromFS(filePath)
	if err != nil {
		switch err {
		case ErrFileNotFound:
			render.Render(w, r, response.ErrNotFound())
		default:
			render.Render(w, r, response.ErrInternalServerError(err))
		}
		log.Err(err).
			Str(logging.FileNameFieldName, filePath).
			Msg("Tried to get file")
		return
	}

	log.Debug().
		Str(logging.FileNameFieldName, filePath).
		Int(logging.FileSizeFieldName, len(data)).
		Msg("File found")
	response.RespondRaw(w, r, data)
}

func HeadFile(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.ErrNotImplementedYet())
}

func PutFile(w http.ResponseWriter, r *http.Request) {
	filePath := chi.URLParam(r, "filepath")
	if filePath == "" {
		render.Render(w, r, response.ErrInternalServerErrorWithCustomMessage("no file name given"))
		return
	}

	if r.ContentLength == 0 {
		render.Render(w, r, response.ErrInternalServerErrorWithCustomMessage("empty body"))
		return
	}

	body := r.Body

	// TODO: Is it advisable to read in the whole file? For now probably not a problem but maybe for larger files
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Err(err).Msg("Tried to read body")
		render.Render(w, r, response.ErrInternalServerError(err))
		return
	}
	err = body.Close() // close as we don't need it anymore
	if err != nil {
		log.Err(err).Msg("Tried to close body")
		render.Render(w, r, response.ErrInternalServerError(err))
		return
	}

	log.Debug().
		Str(logging.FileNameFieldName, filePath).
		Int(logging.FileSizeFieldName, len(data)).
		Msg("Received new file")
	render.Render(w, r, response.ErrNotImplementedYet())
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.ErrNotImplementedYet())
}
