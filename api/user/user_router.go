package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func SetupUserRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/me", GetUserInfo)
	return r
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil || user == "" {
		log.Printf("User not found in context")
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"username": user,
	}
	// Convert map to JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling json:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	_, err = w.Write(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
