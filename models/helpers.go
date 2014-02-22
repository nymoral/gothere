package models

import (
    "fmt"
)

type GuessWithNames struct {
    Team1   string
    Team2   string
    Result1 interface{}
    Result2 interface{}
}

func (g *GuessWithNames) Name() (string) {
    return fmt.Sprintf("%s-%s", g.Team1, g.Team2)
}

func (g *GuessWithNames) Result() (string) {
    if g.Result1 != nil {
        return fmt.Sprintf("%d - %d", g.Result1, g.Result2)
    } else {
        return "-"
    }
}


type GuessContext struct {
    OpenGames   []Game
    Error       bool
    Guesses     []GuessWithNames

}
