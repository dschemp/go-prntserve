package response

import (
	"errors"
	"fmt"
	"github.com/dschemp/go-prntserve/internal/logging"
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
	"net/http"
)

var (
	ErrNoBearerToken      = errors.New("no bearer token passed")
	ErrInvalidBearerToken = errors.New("invalid bearer token")
)

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

func BearerToken(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := middleware.GetReqID(r.Context())
			bearerToken := r.Header.Get("Authorization")
			if len(bearerToken) == 0 {
				bearerTokenUnauthorized(w, r, ErrNoBearerToken, requestID)
				return
			}

			expected := fmt.Sprintf("Bearer %s", token)
			if bearerToken != expected {
				bearerTokenUnauthorized(w, r, ErrInvalidBearerToken, requestID)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func bearerTokenUnauthorized(w http.ResponseWriter, r *http.Request, err error, requestID string) {
	log.Err(err).
		Str(logging.RequestIDFieldName, requestID).
		Msg("Authorization header is empty")
	err = render.Render(w, r, response.ErrUnauthorized(err))
	if err != nil {
		log.Err(err).
			Str(logging.RequestIDFieldName, requestID).
			Msg("Unauthorized message could not be rendered")
	}
}
