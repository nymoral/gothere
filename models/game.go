package models

import (
    "time"
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
    StartsStr   string
}


