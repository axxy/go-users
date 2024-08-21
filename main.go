package main

import (
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	http.HandleFunc("/api/health", health)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"status": "ok"}`)
}
