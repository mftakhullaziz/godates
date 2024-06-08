package common

import (
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteJSONResponse(w, http.StatusUnauthorized, "Missing Authorization header", map[string]string{
				"message": "Missing Authorization header",
			}, 1)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			WriteJSONResponse(w, http.StatusUnauthorized, "Invalid Authorization header format", map[string]string{
				"message": "Invalid Authorization header format",
			}, 1)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Here you can add additional validation of the token, e.g., verifying JWT
		// For now, we'll just pass the token along
		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
