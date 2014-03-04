package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/models"
)

func GetGames(db *sql.DB) ([]models.Game) {
    // Fetches all the games from the db as original models.
    rows, err := db.Query(qGetGames)

    if err != nil {
        log.Fatal(err)
    }

    var G models.Game
    games := make([]models.Game, 0)

    for rows.Next() {
        err := rows.Scan(&G.Team1, &G.Team2, &G.Result1, &G.Result2, &G.StartsStr)
        if err == nil {
            games = append(games, G)
        } else {
            log.Println(err)
        }
    }
    return games
}

func GetUsers(db *sql.DB) ([]models.User) {
    // Generates a list of all the users (not admin).
    // It will be in the left-hand side of main table.
    rows, err := db.Query(qGetUsers)

    if err != nil {
        log.Fatal(err)
    }

    var U models.User
    users := make([]models.User, 0)
    place := 1
    // Calculating place out of db.

    for rows.Next() {
        err := rows.Scan(&U.Firstname, &U.Lastname)
        U.Place = place
        place += 1

        if err == nil {
            users = append(users, U)
        } else {
            log.Println(err)
        }
    }
    return users
}

//func GetGuesses(db *sql.DB) ([]models