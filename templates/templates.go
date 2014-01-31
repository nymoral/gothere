package templates

import (
    "html/template"
    "gothere/config"
    "log"
    "io"
)

func Render(wr io.Writer, name string, data interface{}) (error) {
    dir := config.TemplateDir
    t, err := template.ParseFiles(dir+name+".html")
    if err != nil {
        log.Fatal(err)
    }

    return t.Execute(wr, data)
}
