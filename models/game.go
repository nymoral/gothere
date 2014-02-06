package models

import (
    "time"
    "fmt"
)

type Game struct {
    // Game model.
    // SQL statement in database/SQL/games.sql

    Pk          int
    Team1       string
    Team2       string
    Result1     int
    Result2     int
    Closed      bool
    Happened    bool

    Starts      time.Time
}

func (g *Game) ShortTime() (string) {
    return fmt.Sprintf("%d-%d", g.Starts.Month(), g.Starts.Day())
}

func (g *Game) ListItem() (string) {
    return fmt.Sprintf("<option value=\"%d\">%s %s : %s</option>", g.Pk, g.ShortTime(), g.Team1, g.Team2)
}
