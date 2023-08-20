package middlewares

import (
	"context"
	"net/http"
	"strings"
)

func ExtractOidcToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		oidcToken := authHeader[1]
		ctx := context.WithValue(r.Context(), "oidc", oidcToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
