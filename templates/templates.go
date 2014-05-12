package templates

import (
    "io"
    "log"
    "html/template"
    "gothere/config"
)

var dynamicTemplates = config.Config.DynamicTemplates

var dir = config.Config.TemplateDir

var temps = make(map[string]*template.Template)

func loadTemplates() {
    base := dir + "base.html"

    temps["home"]     = template.Must(template.ParseFiles(dir + "home.html",      base))
    temps["admin"]    = template.Must(template.ParseFiles(dir + "admin.html",     base))
    temps["login"]    = template.Must(template.ParseFiles(dir + "login.html",     base))
    temps["guesses"]  = template.Must(template.ParseFiles(dir + "guesses.html",   base))
    temps["settings"] = template.Must(template.ParseFiles(dir + "settings.html",  base))
    temps["register"] = template.Must(template.ParseFiles(dir + "register.html",  base))
    temps["error"]    = template.Must(template.ParseFiles(dir + "error.html",     base))
    temps["forgot"]   = template.Must(template.ParseFiles(dir + "forgot.html",    base))
    temps["recover"]  = template.Must(template.ParseFiles(dir + "recover.html",   base))

    log.Printf("Templates loaded\n")
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
    t := temps[name]

    err := t.ExecuteTemplate(wr, "base", data)
    if err != nil {
        log.Println(err)
        Render(wr, "error", nil)
    }
}
