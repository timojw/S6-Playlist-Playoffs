package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	// Check if the --development flag is provided
	if len(os.Args) > 1 && os.Args[1] == "--development" {
		log.SetLevel(log.DebugLevel)
		err := godotenv.Load()
		log.Print("Loading .env file")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is game-service!")
}
