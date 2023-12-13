package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"github.com/somatom98/zssn/config"
	"github.com/somatom98/zssn/domain"
	"github.com/somatom98/zssn/inventory"
	"github.com/somatom98/zssn/items"
	"github.com/somatom98/zssn/survivor"
	"go.mongodb.org/mongo-driver/mongo"
)

var conf config.Config
var router *chi.Mux
var mongoDb *mongo.Database

var itemsRepository domain.ItemsRepository
var inventoryRepository domain.InventoryRepository
var survivorRepository domain.SurvivorRepository

var survivorService domain.SurvivorService

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
	itemsRepository = items.NewMockRepository()
	inventoryRepository = inventory.NewMockRepository()
	survivorRepository = survivor.NewMockRepository()

	survivorService = survivor.NewSurvivorService(survivorRepository, inventoryRepository)

	itemsController := items.NewChiController(itemsRepository)
	survivorController := survivor.NewChiController(survivorService)

	router = chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Zombie Survival Social Network!"))
	})
	router.Mount("/items", itemsController.GetRouter())
	router.Mount("/survivors", survivorController.GetRouter())

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
