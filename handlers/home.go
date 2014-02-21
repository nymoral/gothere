package handlers

import (
    "fmt"
    "net/http"
    "gothere/templates"
    "gothere/cookies"
    "gothere/database"
    "gothere/models"
)

type context struct {
    OpenGames   []models.Game
    Error       bool

}

func HomeGet(w http.ResponseWriter, r *http.Request) {
    /*
     * / handler for GET method request.
     * Renders a page only for users with valid sessionid cookie.
     * All the rest are redirected to /login .
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    if username == "" {
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else if is_admin {
        http.Redirect(w, r, "/admin/", http.StatusFound)
    } else {
        var F context
        F.OpenGames = database.GamesList(db, "open")
        F.Error = false
        templates.Render(w, "home", F)
    }

}

func HomePost(w http.ResponseWriter, r *http.Request) {

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    var guess models.Guess
    guess.Userpk = database.GetUserId(db, username)

    var F context
    F.OpenGames = database.GamesList(db, "open")
    F.Error = false

    if username == "" || guess.Userpk < 0{
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        var nr int
        var err error
        //Extract data from request and check if form is valid.
        nr, err = fmt.Sscanf(r.FormValue("result_1"), "%d", &guess.Result1)
        if nr != 1 || err != nil || guess.Result1 < 0 {F.Error = true }
        nr, err = fmt.Sscanf(r.FormValue("result_2"), "%d", &guess.Result2)
        if nr != 1 || err != nil || guess.Result2 < 0 {F.Error = true }
        nr, err = fmt.Sscanf(r.FormValue("game-id"), "%d", &guess.Gamepk)
        if nr != 1 || err != nil {F.Error = true }
        if F.Error {
            templates.Render(w, "home", F)
        } else {
            // Submit a guess.
            database.GiveResult(db, &guess)
            http.Redirect(w, r, "/", http.StatusFound)
        }
    }

}
