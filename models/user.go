package models

import (
    "time"
)

type User struct {
    // User mode.
    // SQL statement in database/SQL/users.sql

    Pk          int

    Email       string
    Password    string

    Firstname   string
    Lastname    string

    Admin       bool
    Joined      time.Time

    Points      int
    Correct     int
}
