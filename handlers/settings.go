package handlers

import (
	"github.com/nymoral/gothere/cookies"
	"github.com/nymoral/gothere/database"
	"github.com/nymoral/gothere/password"
	"github.com/nymoral/gothere/templates"
	"log"
	"net/http"
)

func SettingsGet(w http.ResponseWriter) {
	// /register GET method handler.
	// Just render's the form.

	templates.Render(w, "settings", nil)
}

func SettingsPost(w http.ResponseWriter, r *http.Request) {
	// /settings POST method handler.
	// Validates the form,

	db := database.GetConnection()

	sessionid := cookies.GetCookieVal(r, "sessionid")
	username := cookies.UsernameFromCookie(sessionid)

	if username != "" {
		pass, _ := database.GetPassword(db, username)

		newPassword := r.FormValue("new")
		repeat := r.FormValue("repeat")
		oldPassword := r.FormValue("old")

		if password.Authenticate(oldPassword, pass) && len(newPassword) > 5 && newPassword == repeat {
			hashed := password.NewPassword(newPassword)
			database.ChangePassword(db, username, hashed)
			log.Printf("USER (%s) CHANGED PASSWORD\n", username)
		}
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}
