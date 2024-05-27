package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authenticated, err := ValidateToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Autenticação necessária.")
			return
		}
		if authenticated {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Autenticação necessária.")
		}
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := ValidateToken(r); err != nil {
			return
		}
		next(w, r)
	}
}
