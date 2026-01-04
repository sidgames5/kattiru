package main

import (
	"errors"
	"kattiru/queue"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const version = "v0.1.0"

func main() {
	log.Println("Starting Kattiru " + version)
	defer log.Print("Goodbye!")

	r := mux.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	r.HandleFunc("/queue", queue.HandleDequeue).Methods(http.MethodGet)
	r.HandleFunc("/queue", queue.HandleEnqueue).Methods(http.MethodPost)
	go (func() {
		log.Println("Starting server...")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	})()
}
