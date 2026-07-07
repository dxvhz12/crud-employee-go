package middleware

import (
	"net/http"
	"os"
)

func RequireApiKey(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("API_KEY")

		if apiKey == "" {
			// server salah konfigurasi, bukan salah client
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"server misconfigured: API_KEY not set"}`))
			return
		}

		key := r.Header.Get("X-API-KEY")
		if key != apiKey {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error":"unauthorized"}`))
			return
		}

		next(w, r)
	}
}