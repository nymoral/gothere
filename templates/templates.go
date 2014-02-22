package templates

import (
    "io"
    "html/template"
    "gothere/config"
)

const dir = config.TemplateDir

var homeTemplate, _ = template.ParseFiles(dir + "home.html")

var loginTemplate, _ = template.ParseFiles(dir + "login.html")

var registerTemplate, _ = template.ParseFiles(dir + "register.html")

var adminTemplate, _ = template.ParseFiles(dir + "admin.html")

var errorTemplate, _ = template.ParseFiles(dir + "error.html")

var guessesTemplate, _ = template.ParseFiles(dir + "guesses.html")

func Render(wr io.Writer, name string, data interface{}) {
    /*
     * Renders html template to the response writer.
     */
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
    }

    err := t.Execute(wr, data)
    if err != nil {
        Render(wr, "error", nil)
    }
}
