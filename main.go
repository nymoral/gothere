package main

import (
	"gothere/config"
	"gothere/handlers"
	"log"
	"net/http"
)

/*
 * These functions manage aditional arguments for handlers
 * and request methods.
 */

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

func logout(w http.ResponseWriter, r *http.Request) {
    handlers.Logout(w, r)
}

func errorHand(w http.ResponseWriter) {
    handlers.ErrorGet(w)
}


func main() {
	http.HandleFunc("/",            func(w http.ResponseWriter, r *http.Request) { home(w, r)       })
	http.HandleFunc("/guesses/",    func(w http.ResponseWriter, r *http.Request) { guesses(w, r)    })
	http.HandleFunc("/logout/",     func(w http.ResponseWriter, r *http.Request) { logout(w, r)     })
	http.HandleFunc("/login/",      func(w http.ResponseWriter, r *http.Request) { login(w, r)      })
	http.HandleFunc("/register/",   func(w http.ResponseWriter, r *http.Request) { register(w, r)   })
	http.HandleFunc("/admin/",      func(w http.ResponseWriter, r *http.Request) { admin(w, r)      })
	http.HandleFunc("/error/",      func(w http.ResponseWriter, r *http.Request) { errorHand(w)     })

	if config.ServeStatic {
		// In case go server needs to serve static files.
		// Specified in config file.
		http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(config.Static))))
	}

	log.Println("Server starting")
	http.ListenAndServe(":"+config.Port, nil)
	// Start http server.
}
