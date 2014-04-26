package templates

import (
    "io"
    "log"
    "html/template"
    "gothere/config"
)

const dynamicTemplates = true

var dir = config.Config.TemplateDir

var homeTemplate *template.Template
var loginTemplate *template.Template
var registerTemplate *template.Template
var adminTemplate *template.Template
var errorTemplate *template.Template
var guessesTemplate *template.Template
var settingsTemplate *template.Template


func loadTemplates() {
    homeTemplate, _ = template.ParseFiles(dir + "home.html")
    loginTemplate, _ = template.ParseFiles(dir + "login.html")
    registerTemplate, _ = template.ParseFiles(dir + "register.html")
    adminTemplate, _ = template.ParseFiles(dir + "admin.html")
    errorTemplate, _ = template.ParseFiles(dir + "error.html")
    guessesTemplate, _ = template.ParseFiles(dir + "guesses.html")
    settingsTemplate, _ = template.ParseFiles(dir + "settings.html")
}

func init() {
    if ! dynamicTemplates {
        loadTemplates()
    }
}
func Render(wr io.Writer, name string, data interface{}) {
    // Renders html template to the response writer.
    if dynamicTemplates {
        loadTemplates()
    }
    var t *template.Template

    switch name {
        case "home":
            t = homeTemplate
        case "admin":
            t = adminTemplate
        case "login":
            t = loginTemplate
        case "register":
            t = registerTemplate
        case "error":
            t = errorTemplate
        case "guesses":
            t = guessesTemplate
        case "settings":
            t = settingsTemplate
    }

    err := t.Execute(wr, data)
    if err != nil {
        log.Println(err)
        Render(wr, "error", nil)
    }
}
