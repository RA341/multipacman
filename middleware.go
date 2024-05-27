package main

import (
	"context"
	"database/sql"
	"log"
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
		username := ""
		id := 0
		token, _ := r.Cookie("auth")
		if token == nil {
			log.Print("No auth cookie found")
		} else {
			username, id = auth.GetUserAuthToken(db, token.Value)
		}

		ctx := context.WithValue(r.Context(), "user", username)
		ctx = context.WithValue(ctx, "userId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
