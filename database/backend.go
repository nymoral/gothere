package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/config"
)


func DbInit() (*sql.DB, error) {
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


