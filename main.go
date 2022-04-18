package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/time", time)
	mux.HandleFunc("/health", health)

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Здесь может быть общий код для всех запросов...

	switch r.Method {
	case http.MethodGet:
		// Обработка GET запроса...

	case http.MethodPost:
		// Обработка POST запроса...

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func time(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	var test = "TEST!"
	html := `<doctype html>
        <html>
        <head>
          <title>Hello World</title>
        </head>
        <body>
        <p>
          <a href='/welcome'>Welcome `{{ test }}` </a> | <a href='/message'>Message</a>
        </p>
        </body>
       </html>`
	fmt.Fprintln(w, html)
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusOK)
}
