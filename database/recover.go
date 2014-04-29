package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

func checkDays(days int) (bool) {
    return days <= 2
}

func RecoveryExists(db *sql.DB, pk int) (bool) {
    R := db.QueryRow(qRecoveryExists, pk)
    var days int
    err := R.Scan(&days)

    if err == sql.ErrNoRows {
        return false
    }
    if checkDays(days) {
        return true
    }

    // Is in the db but expired.
    // Can create a new one.
    _, err = db.Exec(qDeleteRecoveryByUser, pk)
    if err != nil {
        log.Fatal(err)
    }
    return false
}

func CheckRecover(db *sql.DB, key string) (bool) {
    R := db.QueryRow(qCheckSql, key)
    var days int
    err := R.Scan(&days)

    if err == sql.ErrNoRows{
        return false
    }
    // Was found
    if err != nil {
        log.Fatal(err)
    }
    if checkDays(days) {
        return true
    }
    _, err = db.Exec(qDeleteRecovery, key)
    if err != nil {
        log.Fatal(err)
    }
    return false
}

func DoRecover(db *sql.DB, key string, password string) {
    r, err := db.Exec(qChangeForKey, password, key)
    if err != nil {
        log.Fatal(err)
    }
    _, err = db.Exec(qDeleteRecovery, key)
    if err != nil {
        log.Fatal(err)
    }
}

func CreateRecovery(db *sql.DB, key string, pk int) {
    _, err := db.Exec(qCreateRecovery, pk, key)
    if err != nil {
        log.Fatal(err)
    }
}

const (
    qCheckSql = "SELECT to_char((NOW() - created)::interval, 'DD') FROM recover WHERE key=$1;"
    qChangeForKey = "UPDATE users SET password=$1 WHERE pk=(SELECT user_pk FROM recover WHERE key=$2);"
    qDeleteRecovery = "DELETE FROM recover WHERE key=$1;"
    qDeleteRecoveryByUser = "DELETE FROM recover WHERE user_pk=$1;"
    qRecoveryExists = "SELECT to_char((NOW() - created)::interval, 'DD') FROM recover WHERE user_pk=$1;"
    qCreateRecovery = "INSERT INTO recover (user_pk, key) VALUES ($1, $2);"
)

