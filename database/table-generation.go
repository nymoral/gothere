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
        err := rows.Scan(&G.Team1, &G.Team2, &G.Result1, &G.Result2, &G.StartsStr, &G.Happened, &G.Closed)
        if err == nil {
            games = append(games, G)
        } else {
            log.Println(err)
        }
    }
    return games
}

func GetUsers(db *sql.DB, pk int) ([]models.User) {
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
        err := rows.Scan(&U.Firstname, &U.Lastname, &U.Pk)
        U.Place = place
        U.LoggedIn = U.Pk == pk
        place += 1

        if err == nil {
            users = append(users, U)
        } else {
            log.Println(err)
        }
    }
    return users
}

func GetGuesses(db *sql.DB, pk int, subsize int, size int) ([][]models.GuessWithPoints) {
    // Returns a slice of models.
    // It will consist of groups of all games for each user.
    rows, err := db.Query(qGetTable, pk)
    if err != nil {
        log.Fatal(err)
    }
    var G models.GuessWithPoints
    guesses := make([][]models.GuessWithPoints, size)

    for j := 0; j < size; j++ {
        guesses[j] = make([]models.GuessWithPoints, subsize)

        for i := 0; i < subsize; i++ {
            rows.Next()
            err := rows.Scan(&G.Result1, &G.Result2, &G.Points, &G.Total, &G.Happened)
            if err == nil {
                guesses[j][i] = G
            } else {
                log.Println(err)
            }
        }
    }
    rows.Close()
    return guesses
}

const (
    qGetGames = "SELECT team1, team2, result1, result2, to_char(starts, 'MM-DD'), happened, closed FROM games ORDER BY starts;"
    qGetUsers = "SELECT firstname, substr(lastname, 1, 1), pk FROM users WHERE admin=false ORDER BY points ASC, correct, pk;"
    qGetTable = "SELECT gs.result1, gs.result2, gs.points, gs.total, G.happened FROM games G LEFT JOIN users AS U ON U.admin=false LEFT JOIN guesses AS gs ON gs.game_pk=G.pk AND gs.user_pk=U.pk AND (U.pk=$1 OR G.closed=true OR G.happened=true) ORDER BY U.points ASC, U.correct, U.pk, G.starts;"
)
