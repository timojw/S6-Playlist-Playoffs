package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/timojw/S6-Playlist-Playoffs/services/voting/internal/kafka"
)

type Vote struct {
	GameID string `json:"game_id"`
	SongID int    `json:"song_id"`
	Points int    `json:"points"`
}

type VoteResult struct {
	SongID int `json:"song_id"`
	Points int `json:"points"`
}

type VotesResponse struct {
	Votes []VoteResult `json:"votes"`
}

var votes = map[string]map[int]int{
	"667b67227b64317349aef419": {
		1:  7,
		2:  2,
		3:  10,
		4:  4,
		5:  9,
		6:  1,
		7:  6,
		8:  3,
		9:  8,
		10: 5,
	},
}

func addVotesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to add votes")

	var voteBatch struct {
		Votes []Vote `json:"votes"`
	}

	if err := json.NewDecoder(r.Body).Decode(&voteBatch); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Decoded votes: %+v", voteBatch)

	for _, vote := range voteBatch.Votes {
		if votes[vote.GameID] == nil {
			votes[vote.GameID] = make(map[int]int)
		}
		votes[vote.GameID][vote.SongID] += vote.Points
		log.Printf("Added %d points to song %d in game %s", vote.Points, vote.SongID, vote.GameID)
	}

	log.Println("Successfully added votes")
	w.WriteHeader(http.StatusCreated)
}

func getVotesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get votes")

	vars := mux.Vars(r)
	gameID := vars["game_id"]
	log.Printf("Game ID from request: %s", gameID)

	// Additional Logging
	log.Printf("Current votes map: %+v", votes)

	if gameID == "" {
		log.Println("Game ID is empty")
		http.Error(w, "Invalid game ID", http.StatusBadRequest)
		return
	}

	gameVotes, ok := votes[gameID]
	if !ok {
		log.Printf("No votes found for game ID: %s", gameID)
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var results []VoteResult
	for songID, points := range gameVotes {
		results = append(results, VoteResult{SongID: songID, Points: points})
		log.Printf("Song ID: %d, Points: %d", songID, points)
	}

	response := VotesResponse{Votes: results}
	log.Println("Successfully retrieved votes")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	if os.Getenv("KAFKA_BOOTSTRAP_SERVER") == "" {
		os.Setenv("KAFKA_BOOTSTRAP_SERVER", "kafka:9092")
	}

	go kafka.RunConsumer()

	r := mux.NewRouter()
	r.HandleFunc("/votes", addVotesHandler).Methods("POST")
	r.HandleFunc("/votes/{game_id}", getVotesHandler).Methods("GET")

	log.Println("Starting voting service on :8082")

	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
