package handlers

import (
    "net/http"
    "gothere/cookies"
    "gothere/templates"
    "gothere/database"
    "gothere/password"
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
        pass , _ := database.GetPassword(db, username)

        newPassword := r.FormValue("new")
        repeat := r.FormValue("repeat")
        oldPassword := r.FormValue("old")

        if password.Authenticate(oldPassword, pass) && len(newPassword) > 5 && newPassword == repeat{
            hashed := password.NewPassword(newPassword)
            database.ChangePassword(db, username, hashed)
        }
    }

    http.Redirect(w, r, "/login", http.StatusFound)
}

