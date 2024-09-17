package routes

import (
	"encoding/json"
	"net/http"
)

func AuthRouteHandler(r *http.ServeMux) http.Handler {
	r.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("This is the register page.")
	})
	return r
}
