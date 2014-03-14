package models

import (
    "fmt"
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

    Place       int

    LoggedIn    bool
}

func (u *User) ShortNameFmt() (string) {
    return fmt.Sprintf("%s %s.", u.Firstname, u.Lastname)
}

func (u *User) Style() (string) {
    if u.LoggedIn {
        return "user"
    }
    return "default"
}

func (u *User) PlaceStyle() (string) {
    if u.LoggedIn {
        return "user_place"
    }
    return "outer place"
}

func (u *User) NameStyle() (string) {
    if u.LoggedIn {
        return "user_name"
    }
    return "default_name"
}
