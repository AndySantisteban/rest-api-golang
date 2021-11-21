package server

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Only GET method is supported")
		return
	}
	fmt.Fprintf(w, "Server On")
}

func getCities(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", cities)
}
func addCity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
