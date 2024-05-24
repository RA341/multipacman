package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

var skipPaths = map[string]bool{
	"/register":    true,
	"/login":       true,
	"/favicon.ico": true,
}

// DurationMiddleware checks duration for each request
func DurationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", r.Method, r.URL.Path, duration)
	})
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
