package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CorsMiddleware used to setup cors options for API.
func CorsMiddleware() func(_ http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Auth-Key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	})
}
