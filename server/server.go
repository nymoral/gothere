package server

import (
    "log"
	"net/http"
	"gothere/config"
)

func ServerInit() {
	http.HandleFunc("/guesses/",    guesses)
	http.HandleFunc("/logout/",     logout)
	http.HandleFunc("/login/",      login)
	http.HandleFunc("/register/",   register)
	http.HandleFunc("/admin/",      admin)
	http.HandleFunc("/error/",      errorHand)
	http.HandleFunc("/settings/",   settings)
	http.HandleFunc("/forgot/",     forgot)
	http.HandleFunc("/",            home)

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

