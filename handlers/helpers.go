package handlers

import (
    "log"
    "net/http"
    "gothere/cookies"
    "gothere/database"
)

func AutoGameClose() {
    db := database.GetConnection()
    toClose := database.GamesToClose(db)
    if len(toClose) > 0 {
        for _, pk := range toClose {
            database.CloseGame(db, pk)
            log.Printf("AUTO CLOSED %s\n", pk)
        }
    }
}

func ChangeSize(w http.ResponseWriter, r *http.Request) {
    tablesize := cookies.GetCookieVal(r, "tablesize")
    if tablesize == "small" {
        cookies.SetCookieVal(w, "tablesize", "full", "/")
    } else {
        cookies.SetCookieVal(w, "tablesize", "small", "/")
    }
    http.Redirect(w, r, "/", http.StatusFound)
}
