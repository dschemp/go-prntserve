package response

import (
	"fmt"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"net/http"
)

// defaultRouter returns a router setup with all kinds of useful middlewares attached.
//
// Included are:
// 	* middleware.RequestID - assigns each request a unique ID
// 	* middleware.Logger - log some useful information to the console
// 	* middleware.Recoverer - handle panics as 500 Server Error
// 	* middleware.RedirectSlashes - removes trailing slashes
//	* middleware.CleanPath - cleans up redundant slashes in request paths
//	* middleware.RealIP - [on demand] uses the IP address provided via the headers for internal processes
//	* render.SetContentType - enables the automatic rendering of output into JSON
func defaultRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.CleanPath)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	if cmd.BehindReverseProxy() {
		r.Use(middleware.RealIP)
	}

	return r
}

func setupRoutes(r chi.Router) {
	r.Get("/{filename}", handler.GetFile)
	r.Head("/{filename}", handler.HeadFile)
	r.With(BearerToken(cmd.Token())).Put("/{filename}", handler.PutFile)
	r.With(BearerToken(cmd.Token())).Delete("/{filename}", handler.DeleteFile)
}

func StartServer() {
	r := defaultRouter()
	setupRoutes(r)

	listenAddress := fmt.Sprintf(":%d", cmd.Port())
	log.Info().
		Str("listenAddress", listenAddress).
		Msg("Starting server")
	log.Printf("Starting listening on %s\n", listenAddress)
	err := http.ListenAndServe(listenAddress, r)
	if err != nil {
		log.Fatal().Err(err)
	}
}
