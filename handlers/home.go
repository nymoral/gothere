package handlers

import (
    "gothere/templates"
    "gothere/cookies"
    "net/http"
    "log"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username == "" {
        // Strange format of a cookie. Gorilla can't decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        templates.Render(w, "home", username)
    }

}
