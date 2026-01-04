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
		Handler: middleware(r),
	}
	r.HandleFunc("/queue", queue.HandleDequeue).Methods(http.MethodGet)
	r.HandleFunc("/queue", queue.HandleEnqueue).Methods(http.MethodPost)

	log.Println("Starting server...")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)
		log.Printf("%s %s %d", r.Method, r.URL.Path, wrappedWriter.statusCode)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
