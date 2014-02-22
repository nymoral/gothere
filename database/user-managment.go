package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/models"
)


func CreateUser(db *sql.DB, user *models.User) {
    /*
     * Adds a user to the database.
     * Assumes that model is correct
     * and has required fields.
     */

    _, err := db.Exec("INSERT INTO users (email, password, firstname, lastname) "+
                        "VALUES ($1, $2, $3, $4);", user.Email, user.Password,
                                                    user.Firstname, user.Lastname)
    if err != nil {
        // Need to check connection to DB.
        log.Fatal(err)
    }
}

func GetPassword(db *sql.DB, email string) (string, bool) {
    /*
     * Fetches hashed users password from the DB.
     * Used to check if user is in the db
     * and for validation / authentication.
     */

    var password string
    var is_admin bool
    R := db.QueryRow("SELECT password, admin FROM users WHERE email=$1;", email)
    err := R.Scan(&password, &is_admin)
    if err == sql.ErrNoRows{
        return "", false
    }
    return password, is_admin
}

func GetUserId(db *sql.DB, username string) (int) {
    // Fetches users pk from the dm.

    var pk int
    R := db.QueryRow("SELECT pk FROM users WHERE email=$1;", username)
    err := R.Scan(&pk)
    if err == sql.ErrNoRows {
        // Usualy not found.
        return -1
    }
    return pk
}

func GetPkAdmin(db *sql.DB, username string) (int, bool) {
    var pk int
    var admin bool
    R := db.QueryRow("SELECT pk, admin FROM users WHERE email=$1;", username)
    err := R.Scan(&pk, &admin)
    if err == sql.ErrNoRows {
        // Usualy not found.
        return -1, false
    }
    return pk, admin
}

