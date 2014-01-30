package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
)


func DbInit() (*sql.DB, error) {
    /*
     * Opens a connection to a postgresql databalse
     * and returns a pointer to sql.DB object and error.
     */
    db, err := sql.Open("postgres", "user=root dbname=gothere sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()

    return db, err

}

func DbClose(db *sql.DB) {
    /* 
     * Closes a connectio to a given sql.DB.
     */
    err := db.Close();
    if err != nil {
        log.Fatal(err)
    }
}


