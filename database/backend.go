package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nymoral/gothere/config"
	"log"
)

var dbConnection *sql.DB

// A connection to the db.

func init() {
	dbConnection = dbInit()
	// Establish the connection.
	dbConnection.SetMaxOpenConns(config.Config.MaxConnections)
	dbConnection.SetMaxIdleConns(config.Config.MaxConnections)
	// Some settings
	err := dbConnection.Ping()
	// Test the connection
	if err != nil {
		log.Println("DB connection test failed.")
		log.Fatal(err)
	}
	log.Printf("Starting db connections. Max open/idle connections: %d\n", config.Config.MaxConnections)
}

func GetConnection() *sql.DB {
	// Passes a connection to a handler.
	return dbConnection
}

func dbInit() *sql.DB {
	// Opens a connection to a postgresql databalse
	// and returns a pointer to sql.DB object.

	statement := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.Config.DbUser,
		config.Config.DbPass,
		config.Config.DbIp,
		config.Config.DbName)

	db, err := sql.Open("postgres", statement)
	if err != nil {
		log.Println("Initial connection to the db failed.")
		log.Fatal(err)
	}

	return db
}
