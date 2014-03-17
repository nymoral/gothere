package server

import (
    "log"
	"net/http"
	"gothere/config"
)

func ServerInit() {
	http.HandleFunc("/",            func(w http.ResponseWriter, r *http.Request) { home(w, r)       })
	http.HandleFunc("/guesses/",    func(w http.ResponseWriter, r *http.Request) { guesses(w, r)    })
	http.HandleFunc("/logout/",     func(w http.ResponseWriter, r *http.Request) { logout(w, r)     })
	http.HandleFunc("/login/",      func(w http.ResponseWriter, r *http.Request) { login(w, r)      })
	http.HandleFunc("/register/",   func(w http.ResponseWriter, r *http.Request) { register(w, r)   })
	http.HandleFunc("/admin/",      func(w http.ResponseWriter, r *http.Request) { admin(w, r)      })
	http.HandleFunc("/error/",      func(w http.ResponseWriter, r *http.Request) { errorHand(w)     })

	if config.Config.ServeStatic {
		// In case go server needs to serve static files.
		// Specified in config file.
		http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(config.Config.Static))))
        log.Println("Serving static.")
	}

    log.Printf("Starting HTTP server at port %s\n", config.Config.Port)
	http.ListenAndServe(":"+config.Config.Port, nil)
	// Start http server.
}

