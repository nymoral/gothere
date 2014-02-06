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

func AdminGet(w http.ResponseWriter, r *http.Request) {
    /*
     * /admin GET method handler.
     * Just render's the form.
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username != "admin" {
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        templates.Render(w, "admin", nil)
    }

}

func AdminPost(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    /*
     * /admin POST method handler.
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username == "admin" {
        var game models.Game
        var err error
        game.Team1= r.FormValue("team1")
        game.Team2= r.FormValue("team2")
        game.Starts, err = time.Parse("2006-01-02 15:04", r.FormValue("starts"))

        if err != nil {
            log.Println(err)
        } else {
            database.CreateGame(db, &game)

        }
        templates.Render(w, "admin", true)
    } else {
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
}

