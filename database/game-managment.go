package database

import (
    _ "github.com/lib/pq"
    "database/sql"
    "log"
    "gothere/models"
)


func CreateGame(db *sql.DB, game *models.Game) {
    /*
     * Adds a game to the database.
     * Assumes that model is correct
     * and has required fields.
     * It only should fail if connection to the DB is no more,
     * therefor it log.Fatal()'s.
     */

    _, err := db.Exec("INSERT INTO games (team1, team2, starts) "+
                        "VALUES ($1, $2, $3);", game.Team1, game.Team2, game.Starts)
    if err != nil {
        log.Fatal(err)
    }
}

