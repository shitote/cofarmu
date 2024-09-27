package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type DBService interface {
	Close() error
}

type Database struct {
	db *sql.DB
}

var dbInstance *Database

func LoadEnv() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	fmt.Println("Environment variables successfully loaded.")
	return nil
}

func NewDatabase() (DBService, error) {
	err := LoadEnv()
	if err != nil {
		return nil, fmt.Errorf("could not load the environment variables %s", err)
	}
	var (
		database = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
		username = os.Getenv("DB_USER")
		port     = os.Getenv("DB_PORT")
		host     = os.Getenv("DB_HOST")
		schema   = os.Getenv("DB_SCHEMA")
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", database, password, host, port, username, schema)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Printf("failed to open database connection: %s", err)
	}

	if err = db.Ping(); err == nil {
		log.Printf("Connection to the database %s", database)
		dbInstance = &Database{db: db}

		return dbInstance, nil
	}

	log.Printf("Failed to ping database: %v", err)

	return nil, fmt.Errorf("faild to connect to the database %s", database)
}

func (s *Database) Close() error {
	log.Printf("Disconnect from database: ")
	return s.db.Close()
}
