package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
)

type DBService interface {
	Close() error
}

type Database struct {
	db *sql.DB
}

var (
	database   = os.Getenv("DB_NAME")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USER")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *Database
)

func NewDatabase() (DBService, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", database, "password", "localhost", port, username, schema)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Printf("failed to open database connection: %s", err)
	}

	if err = db.Ping(); err == nil {
		log.Printf("Connection to the database %s", "postgres")
		dbInstance = &Database{db: db}

		return dbInstance, nil
	}

	log.Printf("Failed to ping database: %v", err)

	return nil, fmt.Errorf("Faild to connect to the database")
}

func (s *Database) Close() error {
	log.Printf("Disconnect frm database: %s", database)
	return s.db.Close()
}
