package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/config"
)

var dbConnection *sql.DB
// A connection to the db.

func init() {
    dbConnection = dbInit()
    dbConnection.SetMaxOpenConns(config.Config.MaxConnections)
    dbConnection.SetMaxIdleConns(config.Config.MaxConnections)
    // Establish the connection.
    _, err := dbConnection.Exec("SELECT pk FROM users WHERE email='admin';")
    if err != nil {
        log.Println("Test query failed")
        log.Fatal(err)
    }
    log.Printf("Starting db connections. Max open/idle connections: %d\n", config.Config.MaxConnections)
}

func GetConnection() (*sql.DB) {
    // Passes a connection to a handler.
    return dbConnection
}

func dbInit() (*sql.DB) {
    // Opens a connection to a postgresql databalse
    // and returns a pointer to sql.DB object.

    statement := "postgres://"
    statement += config.Config.DbUser + ":"
    statement += config.Config.DbPass + "@"
    statement += config.Config.DbIp + "/"
    statement += config.Config.DbName + "?"
    statement += "sslmode=disable"

    db, err := sql.Open("postgres", statement)
    if err != nil {
        log.Println("Initial connection to the db failed.")
        log.Fatal(err)
    }

    return db
}

