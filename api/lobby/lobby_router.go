package lobby

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupLobbyRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create-lobby", registerUser)
	r.Get("/lobbies", loginUser)
	return r
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Home Page"))
	if err != nil {
		return
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {

}
