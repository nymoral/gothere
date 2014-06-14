package models

import (
    "fmt"
    "time"
    "gothere/utils"
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
    n1, n2 := utils.GetShortNames(g.Team1, g.Team2)
    return fmt.Sprintf("%s - %s", n1, n2)
}

func (g *Game) ResultDate() (string) {
    if g.Result1 == nil {
    // Need to return date of the game.
    // It hasn't yet taken place.
        return g.StartsStr
    }
    return fmt.Sprintf("%d - %d", g.Result1, g.Result2)
}

func (g *Game) FullName() (string) {
    return fmt.Sprintf("%s - %s", g.Team1, g.Team2)
}
func (g *Game) Style() (string) {
    if g.Happened {
        return "result"
    }
    return "date"
}

func LastGame(games []Game) (int) {
    for i, g := range games {
        if ! g.Happened {
            return i - 1
        }
    }
    return len(games) - 1
}

