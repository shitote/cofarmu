package server

import (
	"fmt"
	"net/http"
)

func StartServer(r *http.ServeMux) {
	http.ListenAndServe(":8080", r)
	fmt.Println("The server is running on port 8080")

}
