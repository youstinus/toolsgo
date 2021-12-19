package middleware

import "net/http"

// AuthKeyMiddleware checks for valid key before handler.
func AuthKeyMiddleware(key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authKey := r.Header.Get("X-Auth-Key")
			if authKey != key && r.RequestURI != "/metrics" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
