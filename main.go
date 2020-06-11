package main

import (
	"encoding/json"
	"net/http"

	"github.com/jeka1985/go-mod-test/models"
)

func main() {
	http.HandleFunc("/", handleIndex)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Greeting{
		Version: "1.0.0",
		Links:   []string{"some", "links", "for", "discover"},
	})
}
