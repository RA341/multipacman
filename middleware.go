package main

import (
	"context"
	"database/sql"
	"net/http"
	auth "server/api/auth"
)

var skipPaths = map[string]bool{
	"/register":    true,
	"/auth":        true,
	"/favicon.ico": true,
}

func AuthMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := r.Cookie("auth")
		username := auth.GetUserAuthToken(db, token.Value)
		// Pass user information to the context
		ctx := context.WithValue(r.Context(), "user", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
