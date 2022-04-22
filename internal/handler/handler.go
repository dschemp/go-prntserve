package handler

import (
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.ErrNotImplementedYet())
}

func HeadFile(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.ErrNotImplementedYet())
}

func PutFile(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "filename")
	if fileName == "" {
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
		render.Render(w, r, response.ErrInternalServerError(err))
		return
	}
	err = body.Close() // close as we don't need it anymore
	if err != nil {
		render.Render(w, r, response.ErrInternalServerError(err))
		return
	}

	log.Printf(`Received new file "%s" with size %dB`, fileName, len(data))
	render.Render(w, r, response.ErrNotImplementedYet())
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.ErrNotImplementedYet())
}
