package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/models"
)


func CreateUser(db *sql.DB, user *models.User) {
    // Adds a user to the database.
    // Assumes that model is correct
    // and has required fields.

    _, err := db.Exec(qCreateUser, user.Email, user.Password,
                                   user.Firstname, user.Lastname)
    if err != nil {
        // Need to check connection to DB.
        log.Fatal(err)
    }
}

func GetPassword(db *sql.DB, email string) (string, bool) {
    // Fetches hashed users password from the DB.
    // Used to check if user is in the db
    // and for validation / authentication.
    // Also returns a bool that shows if a user is an admin.

    var password string
    var is_admin bool

    R := db.QueryRow(qGetPassword, email)
    err := R.Scan(&password, &is_admin)
    if err == sql.ErrNoRows{
        return "", false
    }
    return password, is_admin
}

func GetPkAdmin(db *sql.DB, username string) (int, bool) {
    // Returns users pk by username and his admin status.

    var pk int
    var admin bool

    R := db.QueryRow(qGetPkAdmin, username)

    err := R.Scan(&pk, &admin)
    if err == sql.ErrNoRows {
        // Usualy not found.
        return -1, false
    }
    return pk, admin
}
