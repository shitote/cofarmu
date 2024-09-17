package server

import (
	"net/http"

	"github.com/shitote/coformu/internal/domain/user/port/routes"
)

func RouterHandler(r *http.ServeMux) {
	r.Handle("/auth", routes.AuthRouteHandler(r))
	r.Handle("/profile", routes.ProfileRouteHandler(r))

}
