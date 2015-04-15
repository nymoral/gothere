package handlers

import (
	"github.com/nymoral/gothere/cookies"
	"github.com/nymoral/gothere/database"
	"log"
	"net/http"
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
