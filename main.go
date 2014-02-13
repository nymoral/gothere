package main

import (
    "net/http"
    "log"
    "os"
    "os/signal"
    "database/sql"
    "gothere/database"
    "gothere/config"
    "gothere/handlers"
)

/*
 * These functions manage aditional arguments for handlers
 * and request methods.
 */

func login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        handlers.LoginGet(w)
    } else {
        handlers.LoginPost(w, r, db)
    }
}

func home(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        handlers.HomeGet(w, r, db)
    }
}

func register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        handlers.RegisterGet(w)
    } else {
        handlers.RegisterPost(w, r, db)
    }
}

func admin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        handlers.AdminGet(w, r, db)
    } else {
        handlers.AdminPost(w, r, db)
    }
}

func main() {

    db, err := database.DbInit()
    // Connection to the db.

    if err != nil {
        // No db - no site.
        log.Fatal(err)
    }

    // Handle ctrl-c
    // to close db connection.
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func(){
        for sig := range c {
            database.DbClose(db)
            _ = sig
            os.Exit(1)
        }
    }()

    http.HandleFunc("/",            func (w http.ResponseWriter, r *http.Request) {home(w, r, db)} )
    http.HandleFunc("/logout/",     func (w http.ResponseWriter, r *http.Request) {handlers.Logout(w, r)} )
    http.HandleFunc("/login/",      func (w http.ResponseWriter, r *http.Request) {login(w, r, db)} )
    http.HandleFunc("/register/",   func (w http.ResponseWriter, r *http.Request) {register(w, r, db)} )
    http.HandleFunc("/admin/",      func (w http.ResponseWriter, r *http.Request) {admin(w, r, db)} )
    http.HandleFunc("/error/",      func (w http.ResponseWriter, r *http.Request) {handlers.ErrorGet(w)} )
    if config.ServeStatic {
        // In case go server needs to serve static files.
        // Specified in config file.
        http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(config.Static))))
    }

    log.Println("Server starting")
    http.ListenAndServe(":" + config.Port, nil)
    // Start http server.
}
