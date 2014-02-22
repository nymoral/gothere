package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/models"
)

// log.Fatal only when executing queries.
// If error occurs need to check connection to DB.

func CreateGame(db *sql.DB, game *models.Game) {
    /*
     * Adds a game to the database.
     * Assumes that model is correct
     * and has required fields.
     */

    _, err := db.Exec("INSERT INTO games (team1, team2, starts) "+
                        "VALUES ($1, $2, $3);", game.Team1, game.Team2, game.Starts)
    if err != nil {
        log.Fatal(err)
    }
}

func OpenGames(db *sql.DB) (*sql.Rows) {
    rows, err := db.Query("SELECT pk, team1, team2, to_char(starts, 'MM-DD') from games "+
                            "WHERE closed=false " +
                            "ORDER BY starts;")
    if err != nil {
        log.Fatal(err)
    }
    return rows

}

func ToFinish(db *sql.DB) (*sql.Rows) {
    // Returns *Rows to be scaned.
    // Games that are to be finished.
    rows, err := db.Query("SELECT pk, team1, team2, to_char(starts, 'MM-DD') from games "+
                            "WHERE happened=false AND closed=true " +
                            "ORDER BY starts;")
    if err != nil {
        log.Fatal(err)
    }
    return rows
}

func GamesList(db *sql.DB, flag string) ([]models.Game) {
    // Returns a slice of models.Game objects
    // that were not closed or not finished based on flag.
    // Only Pk, Team1, Team2, StartsStr fields are filled.

    var rows *sql.Rows
    if flag == "open" {
        rows = OpenGames(db)
    } else {
        rows = ToFinish(db)
    }
    games := make([]models.Game, 0)
    var G models.Game
    for rows.Next() {
        err := rows.Scan(&G.Pk, &G.Team1, &G.Team2, &G.StartsStr)
        if err == nil {
            games = append(games, G)
        }
    }
    if err := rows.Err(); err != nil {
        return nil
    } else {
        return games
    }
}

func CloseGame(db *sql.DB, pk string) {
    // Closes a given game.
    _, err := db.Exec("UPDATE games SET closed = TRUE WHERE pk=$1;", pk)
    if err != nil {
        log.Fatal(err)
    }
}

func FinishGame(db *sql.DB, pk string, t1, t2 int) {
    // Finishes a game.
    // Updates result

    _, err := db.Exec("UPDATE games SET closed = TRUE, happened = TRUE, "+
            "result1=$1, result2=$2 WHERE pk=$3;", t1, t2, pk)
    if err != nil {
        log.Fatal(err)
    }
}
