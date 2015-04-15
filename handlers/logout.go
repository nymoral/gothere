package handlers

import (
	"github.com/nymoral/gothere/cookies"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	// Logs user out (resets cookie)
	// and redirects to /login/ .

	cookies.DeleteSessionId(w)
	http.Redirect(w, r, "/login/", http.StatusFound)
}
