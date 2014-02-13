package handlers

import (
    "net/http"
    "database/sql"
    "gothere/templates"
    "gothere/cookies"
    "gothere/database"
)

func HomeGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    /*
     * / handler for GET method request.
     * Renders a page only for users with valid sessionid cookie.
     * All the rest are redirected to /login .
     */

    sessionid := cookies.GetCookieVal(r, "sessionid")
    username := cookies.UsernameFromCookie(sessionid)
    _, is_admin := database.GetPassword(db, username)

    if username == "" {
        // Gorilla failed to decode it.
        http.Redirect(w, r, "/login/", http.StatusFound)
    } else if is_admin {
        http.Redirect(w, r, "/admin/", http.StatusFound)
    } else {
        templates.Render(w, "home", username)
    }

}
