package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/models"
)


func CreateUser(db *sql.DB, user models.User) {
    _, err := db.Exec("INSERT INTO users (email, password, firstname, lastname) "+
                        "VALUES ($1, $2, $3, $4);", user.Email, user.Password,
                                                   user.Firstname, user.Lastname)
    if err != nil {
        log.Fatal(err)
    }


}

func GetPassword(db *sql.DB, email string) (string) {
    var password string
    R := db.QueryRow("SELECT password FROM users WHERE email=$1;", email)
    err := R.Scan(&password)
    if err != nil {
        return ""
    }
    return password
}

