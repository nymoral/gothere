package handlers

import (
    "gothere/templates"
    "io"
    "log"
)

func LoginGet(w io.Writer) {
    err := templates.Render(w, "login", nil)
    if err != nil {
        log.Fatal(err)
    }
}



