package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/models"
)

func GiveResult(db *sql.DB, guess *models.Guess) {
    // Adds a guess to the database.
    // Assumes that model is correct
    // and has required fields.
    // If the guess already exsists and is ok to change it, updates.
    // It only should fail if connection to the DB is no more,
    // therefor it log.Fatal()'s.

    row := db.QueryRow(qCheckGuess, guess.Gamepk, guess.Userpk)
    // Checks if a guess is in the db.
    var pk int
    err := row.Scan(&pk)
    if err == sql.ErrNoRows {
        // Need to insert.
        _, err := db.Exec(qInsertGuess, guess.Userpk, guess.Gamepk, guess.Result1, guess.Result2)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        // Need to update.
        _, err := db.Exec(qUpdateGuess, guess.Result1, guess.Result2, guess.Gamepk, guess.Userpk)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func UsersGuesses(db *sql.DB, pk int) ([]models.GuessWithNames) {
    // Return's a slice of combined game/guess models.
    // It is rendered in guesses template.

    guesses := make([]models.GuessWithNames, 0)
    var G models.GuessWithNames
    rows, err := db.Query(qUsersGuesses, pk)
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        err := rows.Scan(&G.Team1, &G.Team2, &G.Result1, &G.Result2)
        if err == nil {
            guesses = append(guesses, G)
        }
    }
    return guesses
}
