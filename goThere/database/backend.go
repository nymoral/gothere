package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    //"fmt"
    "log"
)

/*
func (u *User) Rep() {
    fmt.Printf("%d :-: %s \t %s\n", u.id, u.username, u.name)
}

func Show(db *sql.DB) {

    prepQ, err := db.Prepare("SELECT key, username, name FROM testusers ORDER BY username;")
    rows, err := prepQ.Query()

    if err != nil {
        log.Fatal(err)
    }
    var U User
    for rows.Next() {
        rows.Scan(&(U.id), &(U.username), &(U.name))
        U.Rep()
    }

}

func Add(db *sql.DB) {
    Un := "field6"
    Nn := "f; DROP TABLE testusers; ("
    _, err := db.Exec("INSERT INTO testusers (username, name) VALUES ($1, $2);", Un, Nn)
    if err != nil {
        log.Fatal(err)
    }

}

func dbTest() {
    Add(db)
    Show(db)
}
*/

func DbInit() (*sql.DB, error) {
    /*
     * Opens a connection to a postgresql databalse
     * and returns a pointer to sql.DB object and error.
     */
    return sql.Open("postgres", "user=root dbname=gothere sslmode=disable")

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


