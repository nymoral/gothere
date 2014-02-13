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
     * It only should fail if connection to the DB is no more,
     * therefor it log.Fatal()'s.
     */

    _, err := db.Exec("INSERT INTO users (email, password, firstname, lastname) "+
                        "VALUES ($1, $2, $3, $4);", user.Email, user.Password,
                                                    user.Firstname, user.Lastname)
    if err != nil {
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
    if err != nil {
        // Usualy not found.
        return "", false
    }
    return password, is_admin
}
