package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/models"
)

func GiveResult(db *sql.DB, guess *models.Guess) {
    /*
     * Adds a guess to the database.
     * Assumes that model is correct
     * and has required fields.
     * If the guess already exsists and is ok to change it, updates.
     * It only should fail if connection to the DB is no more,
     * therefor it log.Fatal()'s.
     */
     log.Printf("%d %d\n", guess.Gamepk, guess.Userpk)

    row := db.QueryRow("SELECT pk FROM guesses WHERE game_pk=$1 AND user_pk=$2;",
                                guess.Gamepk, guess.Userpk)
    var pk int
    err := row.Scan(&pk)
    if err == sql.ErrNoRows {
        // Need to insert.
        _, err := db.Exec("INSERT INTO guesses (user_pk, game_pk, result1, result2) " +
                            "VALUES ($1, $2, $3, $4);", guess.Userpk, guess.Gamepk,
                                                        guess.Result1, guess.Result2)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        // Need to update.
        _, err := db.Exec("UPDATE guesses SET result1=$1, result2=$2, given=now() WHERE " +
                            "game_pk=$3 AND user_pk=$4;", guess.Result1, guess.Result2,
                                                          guess.Gamepk, guess.Userpk)
        if err != nil {
            log.Fatal(err)
        }
    }
}
