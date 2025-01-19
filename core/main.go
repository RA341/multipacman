package main

import (
	connectcors "connectrpc.com/cors"
	"embed"
	"fmt"
	"github.com/RA341/multipacman/api"
	"github.com/RA341/multipacman/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/gorm"
	"io/fs"
	"net/http"
)

//go:embed web
var frontendDir embed.FS

func main() {
	log.Logger = utils.ConsoleLogger()

	utils.InitConfig()
	db := utils.InitDB()

	// middlewares
	//initMiddlewares(router)
	// ws handler
	//initWSHandler(router, db, m)

	if err := setupServer(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}

func setupServer(db *gorm.DB) error {
	router := api.InitHandlers(db)

	// serve frontend dir
	log.Info().Msgf("Setting up ui files")
	router.Handle("/", getFrontendDir())

	baseUrl := fmt.Sprintf(":%s", viper.GetString("server.port"))
	log.Info().Str("Listening on:", baseUrl).Msg("")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowPrivateNetwork: true,
		AllowedMethods:      connectcors.AllowedMethods(),
		AllowedHeaders:      append(connectcors.AllowedHeaders(), "Authorization"),
		ExposedHeaders:      connectcors.ExposedHeaders(),
	})

	// Use h2c to serve HTTP/2 without TLS
	return http.ListenAndServe(
		baseUrl,
		corsHandler.Handler(h2c.NewHandler(router, &http2.Server{})),
	)
}

func getFrontendDir() http.Handler {
	subFS, err := fs.Sub(frontendDir, "web")
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load frontend directory")
	}
	return http.FileServer(http.FS(subFS))
}

//func initWSHandler(router *chi.Mux, db *sql.DB, m *melody.Melody) {
//	router.Route("/ws", func(r chi.Router) {
//		r.Use(func(handler http.Handler) http.Handler {
//			return AuthMiddleware(db, handler)
//		})
//
//		r.Get("/*", func(writer http.ResponseWriter, request *http.Request) {
//			username := request.Context().Value("user")
//			if username == nil || username == "" {
//				log.Printf("User not found in context")
//				http.Redirect(writer, request, "/login", http.StatusFound)
//				return
//			}
//			err := m.HandleRequest(writer, request)
//			if err != nil {
//				log.Printf("Something went wrong with ws handler")
//				log.Printf(err.Error())
//				return
//			}
//		})
//	})
//
//	initMelody(m)
//}

//func initMiddlewares(router *chi.Mux) {
//	router.Use(middleware.RequestID)
//	router.Use(middleware.RealIP)
//	router.Use(middleware.Logger)
//	router.Use(middleware.Compress(5))
//	router.Use(middleware.Recoverer)
//	router.Use(middleware.Timeout(60 * time.Second))
//}
