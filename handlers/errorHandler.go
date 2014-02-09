package handlers

import (
    "net/http"
    "gothere/templates"
)

func ErrorGet(w http.ResponseWriter) {
    /*
     * Static error page
     */

    templates.Render(w, "error", nil)
}
