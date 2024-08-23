package main

import (
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Second,
	}

	mux.HandleFunc("/hello", hello)
	http.HandleFunc("/api/health", health)

	log.Println("Starting server :8080")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Server failed: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")

}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"status": "ok"}`)
}
