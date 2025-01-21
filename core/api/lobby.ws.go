package api

import (
	"github.com/RA341/multipacman/models"
	"github.com/RA341/multipacman/service"
	"github.com/olahol/melody"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

var (
	mu2           = sync.RWMutex{}
	connections2  = map[uint]*melody.Session{}
	updateChannel = make(chan bool)
)

func initLobbyWsHandler(mux *http.ServeMux, authService *service.AuthService, lobbyService *service.LobbyService) {
	m := melody.New()
	InitLobbyWs(m, lobbyService)

	wsHandler := func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err != nil {
			http.Error(w, "WebSocket connection failed", http.StatusInternalServerError)
			return
		}
	}

	mux.Handle("/api/lobbies", AuthMiddleware(authService, http.HandlerFunc(wsHandler)))
	go sendLobbyUpdates(lobbyService)
}

func InitLobbyWs(m *melody.Melody, lobbyService *service.LobbyService) {
	m.HandleConnect(func(session *melody.Session) {
		HandleConnect(session, m, lobbyService)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		HandleMessage(s, m, msg)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		HandleDisconnect(s, m)
	})
}

func HandleConnect(newPlayerSession *melody.Session, m *melody.Melody, lobbyService *service.LobbyService) {
	user := newPlayerSession.Request.Context().Value("user")
	if user == nil {
		log.Error().Msg("No user found in context")
		err := newPlayerSession.CloseWithMsg([]byte("Unable to determine user info"))
		if err != nil {
			log.Error().Err(err).Msg("Unable close connection")
		}
		return
	}

	verifiedUser := user.(*models.User)
	newPlayerSession.Set("user", verifiedUser)

	mu2.Lock()
	defer mu2.Unlock()
	connections2[verifiedUser.ID] = newPlayerSession

	log.Info().
		Str("username", verifiedUser.Username).
		Uint("id", verifiedUser.ID).
		Msg("client connected")

	// initial data
	jsonData, err := lobbyService.GetAndParseLobbies()
	if err != nil {
		return
	}

	err = newPlayerSession.Write(jsonData)
	if err != nil {
		log.Error().Err(err).Msg("Unable to send lobby data")
	}
}

func HandleDisconnect(s *melody.Session, m *melody.Melody) {
	user, status := s.Get("user")
	if status != true {
		log.Warn().Msg("no value associated with user key")
		return
	}

	id, username := user.(*models.User).ID, user.(*models.User).Username
	connection := connections2[id]
	if connection == nil {
		log.Warn().Uint("id", id).Msg("no connection found for id")
		return
	}
	// remove from connection
	delete(connections2, id)

	err := connection.Close()
	if err != nil {
		log.Error().Err(err).Msg("err while close connection")
		return
	}

	log.Info().Str("username", username).Msg("client disconnected")
}

func sendLobbyUpdates(lobby *service.LobbyService) {
	log.Info().Msg("starting channel watcher")

	for msg := range updateChannel {
		log.Info().Bool("message", msg).Msg("received message")

		jsonData, err := lobby.GetAndParseLobbies()
		if err != nil {
			continue
		}

		for _, session := range connections2 {
			err = session.Write(jsonData)
			if err != nil {
				log.Error().Err(err).Msg("Unable to write to session")
			}
		}
	}
}

// HandleMessage no need to handle message for lobbies
func HandleMessage(s *melody.Session, m *melody.Melody, msg []byte) {}
