package handlers

import (
    "net/http"
    "gothere/templates"
    "gothere/cookies"
)

func HomeGet(w http.ResponseWriter, r *http.Request) {
    /*
     * / handler for GET method request.
     * Renders a page only for users with valid sessionid cookie.
     * All the rest are redirected to /login .
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)

    if username == "" {
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else {
        templates.Render(w, "home", username)
    }

}
