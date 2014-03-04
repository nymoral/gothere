package handlers

import (
    "fmt"
    "net/http"
    "gothere/templates"
    "gothere/cookies"
    "gothere/database"
    "gothere/models"
)

func GuessesGet(w http.ResponseWriter, r *http.Request) {
    // /guesses handler for GET method request.
    // Renders a page only for users with valid sessionid cookie.
    // All the rest are redirected to /login .

    db := database.GetConnection()
    defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    pk, is_admin := database.GetPkAdmin(db, username)

    if username == "" || pk == -1 {
        // Gorilla failed to decode it.
        // Or encoded username does not exist in db.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else if is_admin {
        http.Redirect(w, r, "/admin/", http.StatusFound)
    } else {
        // Fetches users guesses from the db and gets data for
        // result submit dropbox.

        var F models.GuessContext
        F.OpenGames = database.GamesList(db, "open")
        F.Guesses = database.UsersGuesses(db, pk)
        F.Error = false
        templates.Render(w, "guesses", F)
    }

}

func GuessesPost(w http.ResponseWriter, r *http.Request) {
    // /guesses POST method.
    // Checks if user trying to submit is in valid.

    db := database.GetConnection()
    defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    var guess models.Guess
    guess.Userpk, _ = database.GetPkAdmin(db, username)

    var F models.GuessContext
    F.OpenGames = database.GamesList(db, "open")
    F.Guesses = database.UsersGuesses(db, guess.Userpk)
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
            templates.Render(w, "guesses", F)
        } else {
            // Submit a guess.
            database.GiveResult(db, &guess)
            http.Redirect(w, r, "/", http.StatusFound)
        }
    }

}