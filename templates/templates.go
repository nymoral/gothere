package templates

import (
    "html/template"
    "gothere/config"
    "log"
    "io"
)

func Render(wr io.Writer, name string, data interface{}) {
    dir := config.TemplateDir
    t, err := template.ParseFiles(dir+name+".html")
    if err != nil {
        log.Fatal(err)
    }

    err = t.Execute(wr, data)
    if err != nil {
        log.Fatal(err)
    }
}
