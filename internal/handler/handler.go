package handler

import (
    "errors"
    "github.com/dschemp/go-prntserve/internal/logging"
    "github.com/dschemp/go-prntserve/internal/response"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/render"
    "github.com/rs/zerolog/log"
    "io/ioutil"
    "net/http"
)

var (
    ErrEmptyBody       = errors.New("empty body")
    ErrNoFileNameFound = errors.New("no file name given")
)

func GetFile(w http.ResponseWriter, r *http.Request) {
    filePath := chi.URLParam(r, "filepath")

    data, err := GetFileFromStorage(filePath)
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
        // This shouldn't really happen.
        panic(ErrNoFileNameFound)
    }

    if r.ContentLength == 0 {
        render.Render(w, r, response.ErrBadRequest(ErrEmptyBody))
        return
    }

    // TODO: check if file exists, if yes, then throw ErrConflict. only if not, write file
    // for now we can ignore this and check, when we are trying to write file

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

    err = SaveFileToStorage(filePath, data)
    errLog := log.Err(err).
        Str(logging.FileNameFieldName, filePath).
        Int(logging.FileSizeFieldName, len(data))

    // TODO: we should check for file conflict earlier
    if errors.Is(err, ErrFileAlreadyExists) {
        errLog.Msg("file already exists")
        render.Render(w, r, response.ErrConflict(err))
        return
    } else if err != nil {
        errLog.Msg("could not write file")
        render.Render(w, r, response.ErrInternalServerError(err))
        return
    }

    render.Render(w, r, response.FileUploadedSuccessfully())
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
    render.Render(w, r, response.ErrNotImplementedYet())
}
