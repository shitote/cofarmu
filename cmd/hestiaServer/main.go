package main

import (
	"log"
	"net/http"

	"github.com/shitote/coformu/internal/database"
	"github.com/shitote/coformu/internal/server"
)

func main() {
	dbService, dbErr := database.NewDatabase()
	if dbErr != nil {
		log.Fatalf("Error initializing database: %v", dbErr)
	}
	defer func(dbService database.DBService) {
		err := dbService.Close()
		if err == nil {
			log.Fatalf("Error closing database: %v", err)
		}

	}(dbService)

	r := http.NewServeMux()
	server.RouterHandler(r)
	server.StartServer(r)
}
