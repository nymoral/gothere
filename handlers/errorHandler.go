package handlers

import (
	"gothere/templates"
	"log"
	"net/http"
)

func ErrorGet(w http.ResponseWriter) {
	//  Static error page
	templates.Render(w, "error", nil)
	log.Println("ERROR WAS RENDERED")
}
