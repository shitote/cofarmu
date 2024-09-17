package routes

import (
	"encoding/json"
	"net/http"
)

func ProfileRouteHandler(r *http.ServeMux) http.Handler {
	r.HandleFunc("GET /profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("This is the User Profile.")
	})
	return r
}
