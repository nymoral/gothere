package models

import (
    "fmt"
    "time"
)

type Game struct {
    // Game model.
    // SQL statement in database/SQL/games.sql

    Pk          int
    Team1       string
    Team2       string
    Result1     interface{}
    Result2     interface{}
    Closed      bool
    Happened    bool

    Starts      time.Time
    StartsStr   string
}

func (g *Game) NameFmt() (string) {
    return fmt.Sprintf("%s - %s", g.Team1, g.Team2)
}

func (g *Game) ResultDate() (string) {
    if g.Result1 == nil {
    // Need to return date of the game.
    // It hasn't yet taken place.
        return g.StartsStr
    } else {
        return fmt.Sprintf("%d - %d", g.Result1, g.Result2)
    }
}


