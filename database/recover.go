package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func checkDays(days int) bool {
	return days <= 2
}

func RecoveryExists(db *sql.DB, pk int) bool {
	// Checks if there is a recovery in the db by the user pk.
	// If the recovery exists and expired, deletes it and returns false.
	// If there is no recovery reutns true.
	// If there is an active recovery returns true.
	R := db.QueryRow(qRecoveryExists, pk)
	var days int
	err := R.Scan(&days)

	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	if checkDays(days) {
		return true
	}

	// Is in the db but expired.
	// Can delete it.
	_, err = db.Exec(qDeleteRecoveryByUser, pk)
	if err != nil {
		log.Fatal(err)
	}
	return false
}

func CheckRecovery(db *sql.DB, key string) bool {
	// Checks if there is a recovery in the db by the key.
	// If the recovery exists and expired, deletes it and returns false.
	// If there is no recovery reutns true.
	// If there is an active recovery returns true.
	R := db.QueryRow(qCheckRecovery, key)
	var days int
	err := R.Scan(&days)

	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	if checkDays(days) {
		return true
	}

	// Is in the db but expired.
	// Can delete it.
	_, err = db.Exec(qDeleteRecovery, key)
	if err != nil {
		log.Fatal(err)
	}
	return false
}
func DoRecover(db *sql.DB, key string, password string) {
	_, err := db.Exec(qChangeForKey, password, key)
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
	qCheckRecovery        = "SELECT to_char((NOW() - created)::interval, 'DD') FROM recover WHERE key=$1;"
	qChangeForKey         = "UPDATE users SET password=$1 WHERE pk=(SELECT user_pk FROM recover WHERE key=$2);"
	qDeleteRecovery       = "DELETE FROM recover WHERE key=$1;"
	qDeleteRecoveryByUser = "DELETE FROM recover WHERE user_pk=$1;"
	qRecoveryExists       = "SELECT to_char((NOW() - created)::interval, 'DD') FROM recover WHERE user_pk=$1;"
	qCreateRecovery       = "INSERT INTO recover (user_pk, key) VALUES ($1, $2);"
)
