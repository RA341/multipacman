package lobby

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"server/entities"
	"strconv"
)

func SetupLobbyRouter(db *sql.DB, lobbies map[int]*entities.LobbyModel) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", func(writer http.ResponseWriter, request *http.Request) {
		registerLobby(db, lobbies, writer, request)
	})

	r.Get("/lobbies", func(writer http.ResponseWriter, request *http.Request) {
		getAllLobbies(db, lobbies, writer, request)
	})

	r.Delete("/remove/{lobbyId}", func(writer http.ResponseWriter, request *http.Request) {
		deleteLobbies(db, writer, request)
	})
	return r
}

func registerLobby(db *sql.DB, lobbies map[int]*entities.LobbyModel, w http.ResponseWriter, r *http.Request) {
	_, id, allowed := checkAuth(r, w)
	if allowed {
		http.Error(w, "Not allowed login", http.StatusUnauthorized)
		return
	}
	type RequestData struct {
		LobbyName string `json:"lobby_name"`
		UserId    int    `json:"uid"`
	}
	// Parse the JSON body
	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	res := createLobby(db, data.LobbyName, id)
	if res == -2 {
		http.Error(w, "Lobby limit reached, please delete existing lobbies", http.StatusMethodNotAllowed)
		return
	}
	if res == -1 {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// add to lobby list
	lobbies[int(res)] = entities.NewLobbyModel()

	_, err = w.Write([]byte("Lobby created"))
	if err != nil {
		return
	}
}

func getAllLobbies(db *sql.DB, lobbyList map[int]*entities.LobbyModel, w http.ResponseWriter, r *http.Request) {
	_, _, allowed := checkAuth(r, w)
	if allowed {
		http.Error(w, "Not allowed login", http.StatusUnauthorized)
		return
	}
	lobbies := retrieveLobbies(db)
	// add to lobby list
	for _, data := range lobbies {
		lob := lobbyList[data.LobbyId]
		if lob == nil {
			http.Error(w, "Error creating lobby", http.StatusInternalServerError)
			return
		}
		data.Joined = lob.CountPLayers()
	}

	jsonData, err := json.Marshal(lobbies)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}

func deleteLobbies(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	_, id, allowed := checkAuth(r, w)
	if allowed {
		http.Error(w, "Not allowed login", http.StatusUnauthorized)
		return
	}

	lobbyId, err := strconv.Atoi(chi.URLParam(r, "lobbyId"))
	if err != nil {
		http.Error(w, "Error parsing lobbyId", http.StatusBadRequest)
		return
	}
	res := deleteLobby(db, lobbyId, id)
	if !res {
		http.Error(w, "Error deleting lobby", http.StatusBadRequest)
		return
	}
	_, err = w.Write([]byte("Lobby deleted"))
	if err != nil {
		return
	}
}

func checkAuth(r *http.Request, w http.ResponseWriter) (string, int, bool) {
	username := r.Context().Value("user")
	userId := r.Context().Value("userId")
	if username == nil || userId == nil || username == "" || userId == "" {
		http.Error(w, "Unauthorized login first", http.StatusUnauthorized)
		return "", 0, true
	}
	return username.(string), userId.(int), false
}
