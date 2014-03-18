
package server

import (
    "log"
	"net/http"
	"gothere/handlers"
    "gothere/config"
)

// These functions manage aditional arguments for handlers
// and request methods.

func logRequest(r *http.Request) {
    if config.Config.Logging {
        log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
	if r.Method == "GET" {
		handlers.LoginGet(w)
	} else {
		handlers.LoginPost(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
	if r.Method == "GET" {
		handlers.HomeGet(w, r)
	} else {
		handlers.HomePost(w, r)
	}
}

func guesses(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
	if r.Method == "GET" {
		handlers.GuessesGet(w, r)
	} else {
		handlers.GuessesPost(w, r)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
	if r.Method == "GET" {
		handlers.RegisterGet(w)
	} else {
		handlers.RegisterPost(w, r)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
	if r.Method == "GET" {
		handlers.AdminGet(w, r)
	} else {
		handlers.AdminPost(w, r)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
    logRequest(r)
    handlers.Logout(w, r)
}

func errorHand(w http.ResponseWriter, r * http.Request) {
    logRequest(r)
    handlers.ErrorGet(w)
}

