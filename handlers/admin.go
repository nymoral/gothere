package handlers

import (
    "time"
    "log"
    "net/http"
    "database/sql"
    "gothere/templates"
    "gothere/models"
    "gothere/database"
    "gothere/cookies"
)

type FormReturn struct {
    CloseF  bool
    EndF    bool
    OpenGames   []models.Game
}

func AdminGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    /*
     * /admin GET method handler.
     * Just render's the form.
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username != "admin" {
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        var F FormReturn
        F.OpenGames = database.OpenGames(db)
        templates.Render(w, "admin", F)
    }

}

func AdminPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    /*
     * /admin POST method handler.
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    option := r.FormValue("sending")
    var F FormReturn
    F.CloseF = false
    F.EndF = false

    if username == "admin" {
        switch option {
        case "addGame" :
            var game models.Game
            var err error
            game.Team1= r.FormValue("team1")
            game.Team2= r.FormValue("team2")
            game.Starts, err = time.Parse("2006-01-02 15:04", r.FormValue("starts"))

            if err != nil {
                log.Println(err)
            } else {
                database.CreateGame(db, &game)
                http.Redirect(w, r, "/admin", http.StatusFound)
            }

        case "close" :
            pk := r.FormValue("close-game-id")
            database.CloseGame(db, pk)
            http.Redirect(w, r, "/admin", http.StatusFound)

        case "end" :
            http.Redirect(w, r, "/admin", http.StatusFound)
        }
    } else {
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
}
