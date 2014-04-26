
package server

import (
	"net/http"
	"gothere/handlers"
)

// These functions manage aditional arguments for handlers
// and request methods.

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.LoginGet(w)
	} else {
		handlers.LoginPost(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.HomeGet(w, r)
	} else {
		handlers.HomePost(w, r)
	}
}

func guesses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GuessesGet(w, r)
	} else {
		handlers.GuessesPost(w, r)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.RegisterGet(w)
	} else {
		handlers.RegisterPost(w, r)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.AdminGet(w, r)
	} else {
		handlers.AdminPost(w, r)
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.SettingsGet(w)
	} else {
		handlers.SettingsPost(w, r)
	}
}
func logout(w http.ResponseWriter, r *http.Request) {
    handlers.Logout(w, r)
}

func errorHand(w http.ResponseWriter, r * http.Request) {
    handlers.ErrorGet(w)
}
