package handlers

import (
    "log"
    "net/http"
    "gothere/cookies"
    "gothere/database"
    "gothere/password"
    "gothere/templates"
)

func LoginGet(w http.ResponseWriter) {
    // /login handler for GET request.
    // Just renders blank form.

    templates.Render(w, "login", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
    // /login handler for POST request.
    // Tries to validate user.
    // If email / password is OK,
    // new sessionid cookie is set and user is redirected to / .

    db := database.GetConnection()

    username := r.FormValue("username")
    pass := r.FormValue("password")
    remember := r.FormValue("remember") == "1"
    hashed, _ := database.GetPassword(db, username)

    if password.Authenticate(pass, hashed) {
        // Valid password.
        sessionid := cookies.GenerateSessionId(username)
        cookies.SetSessionId(w, sessionid, remember)
        http.Redirect(w, r, "/", http.StatusFound)
        log.Printf("LOGGED IN (%s)\n", username)
    } else {
        templates.Render(w, "login", username)
    }
}
