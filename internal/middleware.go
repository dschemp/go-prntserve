package response

import (
	"errors"
	"fmt"
	"github.com/dschemp/go-prntserve/internal/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

var (
	ErrNoBearerToken      = errors.New("no bearer token passed")
	ErrInvalidBearerToken = errors.New("invalid bearer token")
)

func BearerToken(token string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := middleware.GetReqID(r.Context())
			bearerToken := r.Header.Get("Authorization")
			if len(bearerToken) == 0 {
				err := ErrNoBearerToken
				log.Printf("[%s] %s\n", requestID, err)
				err = render.Render(w, r, response.ErrUnauthorized(err))
				if err != nil {
					log.Println(err)
				}
				return
			}

			expected := fmt.Sprintf("Bearer %s", token)
			if bearerToken != expected {
				err := ErrInvalidBearerToken
				log.Printf("[%s] %s\n", requestID, err)
				err = render.Render(w, r, response.ErrUnauthorized(err))
				if err != nil {
					log.Println(err)
				}
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
