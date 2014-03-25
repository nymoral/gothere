package models

import (
    "fmt"
    "time"
    "gothere/utils"
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

func (user *User) UserValidate(repeat string) (bool) {
    // Registration form validation.
    // Returns true/fales based on weather the form fits
    // requirements.

    if user.Password != repeat {
        // Password don't math.
        return false
    }
    if len(repeat) < 6{
        // Password lenght.
        return false
    }
    if len(user.Firstname) < 1 || len(user.Lastname) < 1 || len(user.Firstname) > 20 || len(user.Lastname) > 30 {
        // To check if not empty and fits in the db.
        return false
    }

    if ! utils.EmailValidation(user.Email) || len(user.Email) > 50 {
        // To check if not empty and fits in the db.
        return false
    }
    return true
}
