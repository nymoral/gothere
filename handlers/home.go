package handlers

import (
    "net/http"
    "database/sql"
    "gothere/models"
    "gothere/config"
    "gothere/cookies"
    "gothere/database"
    "gothere/templates"
)

var _front int
var _back int

func init() {
    _front = config.Config.ShowClosedNr
    _back = config.Config.ShowOpenNr
}


func drawFull(w http.ResponseWriter, db *sql.DB, pk int) {
    var context models.HomeContext
    context.Games = database.GetGames(db)
    lastGame := models.LastGame(context.Games)
    context.Users = database.GetUsers(db, pk)
    context.GamesNr = len(context.Games)
    context.UsersNr = len(context.Users)
    context.Guesses = database.GetGuesses(db, pk, context.GamesNr, context.UsersNr, lastGame)
    templates.Render(w, "home", context)
}

func getSlice(total int, last int) (int, int) {
    if total <= _front + _back {
        return 0, total
    }
    front := last + 1
    back := total - front
    if front >=  _front && back >= _back {
        return last - _front + 1, last + _back + 1
    }
    if front < _front && back >= _back {
        return 0, _front + _back
    }
    return total - (_front + _back), total
}

func drawSmall(w http.ResponseWriter, db *sql.DB, pk int) {
    var context models.HomeContext
    context.Users = database.GetUsers(db, pk)
    context.UsersNr = len(context.Users)

    allGames := database.GetGames(db)
    lastGame := models.LastGame(allGames)
    s, e := getSlice(len(allGames), lastGame)
    context.Games = allGames[s:e]
    context.Guesses = database.GetSmall(db, pk, context.UsersNr, s, e - s, lastGame)
    templates.Render(w, "small", context)
}

func HomeGet(w http.ResponseWriter, r *http.Request) {
    // / handler for GET method request.
    // Renders a page only for users with valid sessionid cookie.
    // All the rest are redirected to /login .

    db := database.GetConnection()

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    tablesize := cookies.GetCookieVal(r, "tablesize")
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
        if tablesize == "small" {
            drawSmall(w, db, pk)
        } else {
            drawFull(w, db, pk)
        }

    }
}

func HomePost(w http.ResponseWriter, r *http.Request) {
    // /home POST handler.

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username == "" {
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        http.Redirect(w, r, "/", http.StatusFound)
    }
}
