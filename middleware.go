package main

import (
	"context"
	"net/http"
)

var skipPaths = map[string]bool{
	"/register":    true,
	"/login":       true,
	"/favicon.ico": true,
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.Header.Get("Authorization")
		// Example: Verify the token here and extract user information
		// For simplicity, let's assume we have a function ValidateToken
		// which returns user information if the token is valid.

		// Pass user information to the context
		ctx := context.WithValue(r.Context(), "user", nil)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
