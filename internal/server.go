package response

import (
	"fmt"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"github.com/dschemp/go-prntserve/internal/logging"
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
//  * jwtauth.Verifier - checks the specified JWT token in the request
func defaultRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.CleanPath)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(jwtauth.Verifier(cmd.JWTTokenAuth()))

	if cmd.BehindReverseProxy() {
		r.Use(middleware.RealIP)
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, response.ErrNotFound())
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, response.ErrMethodNotAllowed())
	})

	return r
}

// setupRoutes assigns all routes for this application to a specified chi.Router.
func setupRoutes(r chi.Router) {
	// For now only support files on top-most level and not directories.
	// Implementing a directory server (like a file server) is more complex.
	// See e. g. https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go#L53
	fileRoute := "/{filepath:[a-zA-Z0-9-_.]+}"
	r.Get(fileRoute, handler.GetFile)
	r.Head(fileRoute, handler.HeadFile)
	r.With(JWTAuthenticator).Put(fileRoute, handler.PutFile)
	r.With(JWTAuthenticator).Delete(fileRoute, handler.DeleteFile)
}

// createInitialAccessJWT creates a JWT token (string) with no payload. This is intended to be used as an initial authentication token.
func createInitialAccessJWT() string {
	_, token, _ := cmd.JWTTokenAuth().Encode(nil)
	return token
}

// StartServer sets up everything needed to run the server and then runs it.
func StartServer() {
	r := defaultRouter()
	setupRoutes(r)

	token := createInitialAccessJWT()
	log.Info().
		Str("access_token", token).
		Msg("Created initial access token")

	listenAddress := fmt.Sprintf(":%d", cmd.Port())
	log.Info().
		Str(logging.ListenAddressFieldName, listenAddress).
		Msg("Starting server")
	err := http.ListenAndServe(listenAddress, r)
	if err != nil {
		log.Fatal().Err(err)
	}
}
