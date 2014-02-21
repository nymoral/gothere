package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/config"
)

func DbInit() (*sql.DB) {
    /*
     * Opens a connection to a postgresql databalse
     * and returns a pointer to sql.DB object and error.
     */

    uname := " user=" + config.DbUser
    dname := " dbname=" + config.DbName

    var pass string

    if config.DbPass != "" {
        pass = " password=" + config.DbPass
    } else {
        pass = ""
    }

    openStatement := "sslmode=disable" + dname + uname + pass
    db, err := sql.Open("postgres", openStatement)
    if err != nil {
        log.Fatal(err)
    }

    return db
}
