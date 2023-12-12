package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"github.com/somatom98/zssn/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var conf config.Config
var router *chi.Mux
var mongoDb *mongo.Database

func init() {
	conf, err := config.GetFromYaml()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to load config")
	}
	log.Info().
		Msg("Config loaded")

	mongoDb, err = conf.GetMongoDb(context.Background())
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to MongoDB")
	}
	log.Info().Msg("MongoDB connected")
}

func main() {
	router = chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Zombie Survival Social Network!"))
	})

	log.Info().
		Msg("HTTP Server starting")
	httpSrv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
	}
	httpSrv.Addr = ":8080"
	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start HTTP server")
	}
}
