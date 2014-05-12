package handlers

import (
    "log"
    "net/http"
    "gothere/templates"
)

func ErrorGet(w http.ResponseWriter) {
    //  Static error page
    templates.Render(w, "error", nil)
    log.Println("ERROR WAS RENDERED")
}
