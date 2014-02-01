package main

import (
    //"gothere/tests"
    "net/http"
    "gothere/config"
    "gothere/handlers"
    "log"
    "gothere/database"
    "database/sql"
)

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
        // TODO
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
}

func register(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    if r.Method == "GET" {
        handlers.RegisterGet(w)
    } else {
        handlers.RegisterPost(w, r, db)
    }

}

func main() {
    //tests.Test()

    db, err := database.DbInit()

    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", home)
    http.HandleFunc("/login/", login)
    http.HandleFunc("/register/", func (w http.ResponseWriter, r *http.Request){ register(w, r, db)})
    if config.ServeStatic {
        http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(config.Static))))
    }
    log.Println("Server starting")
    http.ListenAndServe(":" + config.Port, nil)
}


