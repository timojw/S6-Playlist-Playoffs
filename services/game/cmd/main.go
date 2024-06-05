// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	log "github.com/sirupsen/logrus"

// 	"github.com/joho/godotenv"
// )

// type Game struct {
// 	Name     string    `json:"name"`
// 	Deadline time.Time `json:"deadline"`
// }

// func createGameHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var game Game
// 	err := json.NewDecoder(r.Body).Decode(&game)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	log.Printf("Game created: %+v\n", game)
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(game)
// }

// func main() {
// 	if len(os.Args) > 1 && os.Args[1] == "--development" {
// 		log.SetLevel(log.DebugLevel)
// 		err := godotenv.Load()
// 		log.Print("Loading .env file")
// 		if err != nil {
// 			log.Fatal("Error loading .env file")
// 		}
// 	}
// 	http.HandleFunc("/game", createGameHandler)
// 	fmt.Println("Starting server at port 8081")
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've hit %s\n", r.URL.Path)
	log.Printf("Served request on %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // Set the handler for the root path

	// Configure logging
	logFile, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Starting server on port 8081...")

	// Start the server on port 8081 and log any errors
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
