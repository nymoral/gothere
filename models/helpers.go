package models

import (
    "fmt"
)

type GuessWithNames struct {
    // A modele for guesses template.
    // It holds game team names
    // and int/nil for (not)given result.

    Team1   string
    Team2   string
    Result1 interface{}
    Result2 interface{}
}

func (g *GuessWithNames) Name() (string) {
    // Name formating
    return fmt.Sprintf("%s-%s", g.Team1, g.Team2)
}

func (g *GuessWithNames) Result() (string) {
    // Result formating
    if g.Result1 != nil {
        return fmt.Sprintf("%d - %d", g.Result1, g.Result2)
    } else {
        return " "
    }
}


type GuessContext struct {
    // Complete model for guess template.
    OpenGames   []Game
    Error       bool
    Guesses     []GuessWithNames

}

type RegisterContext struct {
    Flag        bool
    Firstname   string
    Lastname    string
    Email       string
}

