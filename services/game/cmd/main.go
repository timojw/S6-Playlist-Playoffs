package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/kafka"
)

type Game struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

// hardcoded :(
var games = map[string]Game{
	"1": {ID: "1", Name: "skibidi game", Deadline: "2024-06-12T21:59"},
	"2": {ID: "2", Name: "rara ra", Deadline: "2024-07-12T21:59"},
}

func addGameHandler(w http.ResponseWriter, r *http.Request) {
	var game Game

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// id generate
	game.ID = uuid.New().String()
	games[game.ID] = game

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	game, exists := games[id]
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	gameID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	kafka.RunProducer(gameID, true)

	json.NewEncoder(w).Encode(game)
}

func main() {
	// if os.Getenv("KAFKA_BOOTSTRAP_SERVER") == "" {
	// 	os.Setenv("KAFKA_BOOTSTRAP_SERVER", "kafka:9092")
	// }

	r := mux.NewRouter()
	r.HandleFunc("/game", addGameHandler).Methods("POST")
	r.HandleFunc("/game/{id}", getGameHandler).Methods("GET")

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
