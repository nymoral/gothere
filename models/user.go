package models

import (
    "time"
)

type User struct {
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
