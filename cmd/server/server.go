package server

import (
	"net/http"

	"github.com/sithumonline/movie-box/api/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RunServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start movie-box server",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/movie-box/{name}", handler.AddMovie)
	//	r.Get("/movie-box", handler.GetLogs)

	log.Info("movie-box started listening on 3080")

	if err := http.ListenAndServe(":3080", r); err != nil {
		log.Fatal(err)
	}
}
