package handlers

import (
    "gothere/templates"
    "gothere/cookies"
    "io"
    "net/http"
    "log"
)

func LoginGet(w io.Writer) {
    templates.Render(w, "login", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
    username := r.FormValue("username")
    sessionid := cookies.GenerateSessionId(username)
    cookies.SetSessionId(w, sessionid, false)
    http.Redirect(w, r, "/", 302)
}



