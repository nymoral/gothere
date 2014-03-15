package handlers

import (
    "time"
    "strconv"
    "net/http"
    "gothere/utils/games"
    "gothere/models"
    "gothere/templates"
    "gothere/database"
    "gothere/cookies"
)

func AdminGet(w http.ResponseWriter, r *http.Request) {
    // /admin GET method handler.
    // Just render's the form.

    db := database.GetConnection()
    defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    if ! is_admin {
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        var F models.AdminContext
        F.OpenGames = database.GamesList(db, "open")
        F.NotFinish = database.GamesList(db, "finish")
        templates.Render(w, "admin", F)
    }
}

func AdminPost(w http.ResponseWriter, r *http.Request) {
    // /admin POST method handler.

    db := database.GetConnection()
    defer database.RecycleConnection(db)

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    option := r.FormValue("sending")
    // Each html form has a hidden input which
    // is sent only when a specific form is submited.
    var F models.AdminContext
    F.CloseF = false
    F.EndF = false

    if is_admin {
        switch option {
        case "addGame" :
            // Creating a game and sending it into db.
            var game models.Game
            var err error
            game.Team1= r.FormValue("team1")
            game.Team2= r.FormValue("team2")
            game.Starts, err = time.Parse("2006-01-02 15:04", r.FormValue("starts"))
            // Need to parse a string from request.

            if err != nil || ! games.HasShortName(game.Team1) || ! games.HasShortName(game.Team2) {
                // Checks if teams names are valid.
                http.Redirect(w, r, "/error", http.StatusFound)
            } else {
                database.CreateGame(db, &game)
                http.Redirect(w, r, "/admin", http.StatusFound)
            }

        case "close" :
            // Closes a game. Nothing to check.
            pk := r.FormValue("close-game-id")
            database.CloseGame(db, pk)
            http.Redirect(w, r, "/admin", http.StatusFound)

        case "end" :
            // Finishes a game.
            // TODO Recalculate scores.

            t1 := r.FormValue("team1")
            t2 := r.FormValue("team2")
            pk := r.FormValue("finish-game-id")
            n1, er1 := strconv.Atoi(t1)
            n2, er2 := strconv.Atoi(t2)
            intPk, er3 := strconv.Atoi(pk)
            if er1 != nil || er2 != nil || er3 != nil {
                // Checks form data.
                http.Redirect(w, r, "/error", http.StatusFound)
            } else {
                database.FinishGame(db, pk, n1, n2)
                database.CalcPoints(db, intPk, n1, n2)
                http.Redirect(w, r, "/admin", http.StatusFound)
            }
        }
    } else {
        // Not an admin tried subminting data.
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
}
