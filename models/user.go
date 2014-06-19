package models

import (
	"fmt"
	"gothere/utils"
	"time"
)

type User struct {
	// User mode.
	// SQL statement in database/SQL/users.sql

	Pk int

	Email    string
	Password string

	Firstname string
	Lastname  string

	Admin  bool
	Joined time.Time

	Points  int
	Correct int

	Place int

	LoggedIn bool
}

func (u *User) ShortNameFmt() string {
	suffix := ""
	if len(u.Lastname) > 0 {
		suffix = u.Lastname[:1]
	}
	return fmt.Sprintf("%s %s.", u.Firstname, suffix)
}

func (u *User) Style() string {
	if u.LoggedIn {
		return "user"
	}
	return "default"
}

func (u *User) PlaceStyle() string {
	if u.LoggedIn {
		return "user_place"
	}
	return "outer place"
}

func (u *User) NameStyle() string {
	if u.LoggedIn {
		return "user_name"
	}
	return "default_name"
}

func (user *User) UserValidate(repeat string) string {
	// Registration form validation.
	// Returns true/fales based on weather the form fits
	// requirements.

	if len(user.Firstname) < 1 || len(user.Lastname) < 1 || len(user.Firstname) > 20 || len(user.Lastname) > 30 {
		// To check if not empty and fits in the db.
		return "Vardas turi būti 1-20, pavardė 1-30 simbolių ilgio."
	}
	if !utils.EmailValidation(user.Email) || len(user.Email) > 50 {
		// To check if not empty and fits in the db.
		return "El. pašto adresas turi būti validus ir iki 50 simbolių ilgio"
	}
	if len(repeat) < 6 {
		// Password lenght.
		return "Slatažodis turi būti bent šešių (6) simbolių ilgio"
	}
	if user.Password != repeat {
		// Password don't math.
		return "Slaptažodžiai nesutampa."
	}

	return ""
}
