package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type ViewData struct {
	NewYork string `json:"NewYork"`
	Berlin  string `json:"NewYork"`
	Tokyo   string `json:"NewYork"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Time)
	mux.HandleFunc("/health", health)

	err := http.ListenAndServe(":80", mux)
	log.Fatal(err)
}

func Time(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	now := time.Now()
	loc, _ := time.LoadLocation("UTC")

	loc, _ = time.LoadLocation("America/New_York")
	var New_York = now.In(loc)

	loc, _ = time.LoadLocation("Europe/Berlin")
	var Berlin = now.In(loc)

	loc, _ = time.LoadLocation("Asia/Tokyo")
	var Tokyo = now.In(loc)

	data := ViewData{
		NewYork: New_York.String(),
		Berlin:  Berlin.String(),
		Tokyo:   Tokyo.String(),
	}

	tmpl := template.Must(template.New("data").Parse(`<div>
	    <p>New York: {{.NewYork}} </p>
	    <p>Berlin:  {{.Berlin}} </p>
	    <p>Tokyo:{{.Tokyo}} </p>
	</div>`))
	tmpl.Execute(w, data)

}

func health(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusOK)
}
