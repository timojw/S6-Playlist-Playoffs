package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/routes"
)

type Game = models.Game

// hardcoded :(
var games = map[string]Game{
	"1": {ID: "1", Name: "skibidi game", Deadline: "2024-06-12T21:59"},
	"2": {ID: "2", Name: "rara ra", Deadline: "2024-07-12T21:59"},
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Host = "localhost:" + config.EnvPort()

	// config.SetIndexes()

	routes.GameRoutes(router)
	routes.SwaggerRoutes(router)

	router.Run("0.0.0.0:8081")

	// if os.Getenv("KAFKA_BOOTSTRAP_SERVER") == "" {
	// 	os.Setenv("KAFKA_BOOTSTRAP_SERVER", "kafka:9092")
	// }

	// r := mux.NewRouter()
	// r.HandleFunc("/game", addGameHandler).Methods("POST")
	// r.HandleFunc("/game/{id}", getGameHandler).Methods("GET")

	// log.Println("Starting server on :8081")
	// if err := http.ListenAndServe(":8081", r); err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }
}
