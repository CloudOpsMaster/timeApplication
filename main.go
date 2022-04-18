package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	HandleFunction()
}

func HandleFunction() {

	var port = "8080"
	r := mux.NewRouter()

	r.HandleFunc("/", TestFunc).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	var hello string = "Hello GO"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hello)
}
