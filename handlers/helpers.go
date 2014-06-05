package handlers

import (
    "log"
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
