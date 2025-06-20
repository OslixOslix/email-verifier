package middleware

import (
	"net/http"
	"os"
)

// APIKeyMiddleware checks the 'apikey' header against the API_KEY environment variable.
// If the key doesn't match, it returns 403 Forbidden with a Russian message.
func APIKeyMiddleware(next http.Handler) http.Handler {
	requiredKey := os.Getenv("API_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if requiredKey != "" {
			if r.Header.Get("apikey") != requiredKey {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte("Доступ запрещен"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
