package server

import "net/http"

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCities(w, r)
		case http.MethodPost:
			addCity(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

	})
}
