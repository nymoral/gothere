package main

import (
    //"gothere/tests"
    "net/http"
    "gothere/config"
    "gothere/handlers"
    "fmt"
)

func handler(w http.ResponseWriter, r * http.Request) {
    fmt.Fprintf(w, "Hello, Internet!")

}

func login(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        handlers.LoginGet(w)
    }
}

func main() {
    //tests.Test()

    http.HandleFunc("/", handler)
    http.HandleFunc("/login/", login)
    if config.ServeStatic {
        http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(config.Static))))
    }
    http.ListenAndServe(":" + config.Port, nil)

}


