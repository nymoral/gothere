package server

import (
	"gothere/handlers"
	"net/http"
	"time"
)

// These functions manage aditional arguments for handlers
// and request methods.

func runAuto() {
	for true {
		handlers.AutoGameClose()
		time.Sleep(1 * time.Minute)
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.LoginPost(w, r)
	} else {
		handlers.LoginGet(w)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.HomePost(w, r)
	} else {
		handlers.HomeGet(w, r)
	}
}

func guesses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.GuessesPost(w, r)
	} else {
		handlers.GuessesGet(w, r)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.RegisterPost(w, r)
	} else {
		handlers.RegisterGet(w)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.AdminPost(w, r)
	} else {
		handlers.AdminGet(w, r)
	}
}

func settings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.SettingsPost(w, r)
	} else {
		handlers.SettingsGet(w)
	}
}

func forgot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlers.ForgotPost(w, r)
	} else {
		handlers.ForgotGet(w, r)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	handlers.Logout(w, r)
}

func errorHand(w http.ResponseWriter, r *http.Request) {
	handlers.ErrorGet(w)
}
func changeSize(w http.ResponseWriter, r *http.Request) {
	handlers.ChangeSize(w, r)
}
