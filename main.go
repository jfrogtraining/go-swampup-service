package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var results []string

func main() {
	// use PORT environment variable, or default to 7888
	port := "7888"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()

	server.HandleFunc("/ping", ping)
	server.HandleFunc("/event", event)
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

func event(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message  string    `json:"message"`
		Timestam time.Time `json:"timestamp"`
	}{
		"Welcome to GopherCon 2019",
		time.Now(),
	}

	sendJsonResponse(w, http.StatusOK, response)
}

// ping responds to the request with a plain-text "Ok" message.
func ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	fmt.Fprintf(w, "Pong")
}

func sendJsonResponse(res http.ResponseWriter, httpCode int, payload interface{}) {
	log.Infof("Sending %v response: %#v", httpCode, payload)
	res.Header().Set("Content-Type", "application/json")
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to send response: %s", err)))
	}
	res.WriteHeader(httpCode)
	res.Write(jsonPayload)
}
