package response

import (
	"errors"
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

var (
	ErrNoBearerToken      = errors.New("no bearer token passed")
	ErrInvalidBearerToken = errors.New("invalid bearer token")
)

// JWTAuthenticator is a middleware intended to be used as a gatekeeper for routes which need authentication.
// This DOES NOT handle authorization!
func JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			render.Render(w, r, response.ErrUnauthorized(err))
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			render.Render(w, r, response.ErrUnauthorized(err))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
