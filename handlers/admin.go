package handlers

import (
    "time"
    "strconv"
    "net/http"
    "gothere/templates"
    "gothere/models"
    "gothere/database"
    "gothere/cookies"
)

type FormReturn struct {
    CloseF  bool
    EndF    bool
    OpenGames   []models.Game
    NotFinish   []models.Game
}

func AdminGet(w http.ResponseWriter, r *http.Request) {
    /*
     * /admin GET method handler.
     * Just render's the form.
     */
    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    if ! is_admin {
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        var F FormReturn
        F.OpenGames = database.GamesList(db, "open")
        F.NotFinish = database.GamesList(db, "finish")
        templates.Render(w, "admin", F)
    }
}

func AdminPost(w http.ResponseWriter, r *http.Request) {
    /*
     * /admin POST method handler.
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    option := r.FormValue("sending")
    var F FormReturn
    F.CloseF = false
    F.EndF = false

    if is_admin {
        switch option {
        case "addGame" :
            var game models.Game
            var err error
            game.Team1= r.FormValue("team1")
            game.Team2= r.FormValue("team2")
            game.Starts, err = time.Parse("2006-01-02 15:04", r.FormValue("starts"))

            if err != nil {
                http.Redirect(w, r, "/error", http.StatusFound)
            } else {
                database.CreateGame(db, &game)
                http.Redirect(w, r, "/admin", http.StatusFound)
            }

        case "close" :
            pk := r.FormValue("close-game-id")
            database.CloseGame(db, pk)
            http.Redirect(w, r, "/admin", http.StatusFound)

        case "end" :
            t1 := r.FormValue("team1")
            t2 := r.FormValue("team2")
            pk := r.FormValue("finish-game-id")
            n1, er1 := strconv.Atoi(t1)
            n2, er2 := strconv.Atoi(t2)
            if er1 != nil || er2 != nil {
                http.Redirect(w, r, "/error", http.StatusFound)
            } else {
                database.FinishGame(db, pk, n1, n2)
                http.Redirect(w, r, "/admin", http.StatusFound)
            }
        }
    } else {
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
}
