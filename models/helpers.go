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
    Flag        string
    Firstname   string
    Lastname    string
    Email       string
}

type HomeContext struct {
    Games       []Game
    Users       []User
    Guesses     [][]GuessWithPoints

    GamesNr     int
    UsersNr     int
}


type AdminContext struct {
    CloseF  bool
    EndF    bool
    OpenGames   []Game
    NotFinish   []Game
}

type GuessWithPoints struct {
    Result1     interface{}
    Result2     interface{}

    Points      interface{}
    Total       interface{}

    Happened    bool
}

func (g *GuessWithPoints) ResultFmt() (string) {
    if g.Result1 == nil {
        return " "
    } else {
        if g.Result1.(int64) >= 0 {
            return fmt.Sprintf("%d : %d", g.Result1, g.Result2)
        } else {
            return "-"
        }
    }
}

func (g *GuessWithPoints) PointsFmt() (string) {
    if g.Happened {
        if g.Points == nil {
            return "-"
        } else {
            return fmt.Sprintf("%d / %d", g.Points, g.Total)
        }
    } else {
        return " "
    }
}
func (g *GuessWithPoints) Style() (string) {
    if g.Happened {
        return "happened"
    }
    return "not_happened"
}

type Points struct {
    UserPk  int
    Points  int
    Result1 interface{}
    Result2 interface{}

    Given   bool
    Total   int
}
