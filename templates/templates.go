package templates

import (
    "html/template"
    "log"
    "io"
    "gothere/config"
)

func Render(wr io.Writer, name string, data interface{}) {
    /*
     * Renders html template to the response writer.
     * Given template should exist in ./html as name + ".html"
     */

    dir := config.TemplateDir
    t, err := template.ParseFiles(dir + name + ".html")
    if err != nil {
        log.Fatal(err)
    }

    err = t.Execute(wr, data)
    if err != nil {
        log.Fatal(err)
    }
}
