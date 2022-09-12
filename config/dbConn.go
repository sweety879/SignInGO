package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	controllers "go-psql-gin/controllers"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sweety.seela"
	password = "postgres"
	dbname   = "credentials"
)

// Connecting to db
func Connect() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	controllers.InitiateDB(db)
	log.Println("Successfully connected!")

	return db

}
