package handlers

import (
    "net/http"
    "gothere/models"
    "gothere/templates"
    "gothere/cookies"
    "gothere/database"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
    // / handler for GET method request.
    // Renders a page only for users with valid sessionid cookie.
    // All the rest are redirected to /login .

    db := database.GetConnection()
    defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    pk, is_admin := database.GetPkAdmin(db, username)

    if username == "" || pk == -1 {
        // Gorilla failed to decode it.
        // Or user is not in the db.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else if is_admin {
        // Admin needs to be redirected to
        // administration site.
        http.Redirect(w, r, "/admin/", http.StatusFound)
    } else {
        // Render home.
        var context models.HomeContext
        context.Games = database.GetGames(db)
        context.Users = database.GetUsers(db, pk)
        context.GamesNr = len(context.Games)
        context.UsersNr = len(context.Users)
        context.Guesses = database.GetGuesses(db, pk, context.GamesNr, context.UsersNr)
        templates.Render(w, "home", context)
    }
}

func HomePost(w http.ResponseWriter, r *http.Request) {
    // /home POST handler.
    // TODO :

    //db := database.GetConnection()
    //defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username == "" {
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        http.Redirect(w, r, "/", http.StatusFound)
    }
}
