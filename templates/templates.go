package templates

import (
	"fmt"
	"github.com/nymoral/gothere/config"
	"html/template"
	"io"
	"log"
)

var dynamicTemplates = config.Config.DynamicTemplates

var dir = config.Config.TemplateDir

var temps = make(map[string]*template.Template)

func loadTemplates() {
	base := dir + "base.html"

	temps["home"] = template.Must(template.ParseFiles(dir+"home.html", base))
	temps["admin"] = template.Must(template.ParseFiles(dir+"admin.html", base))
	temps["login"] = template.Must(template.ParseFiles(dir+"login.html", base))
	temps["guesses"] = template.Must(template.ParseFiles(dir+"guesses.html", base))
	temps["settings"] = template.Must(template.ParseFiles(dir+"settings.html", base))
	temps["register"] = template.Must(template.ParseFiles(dir+"register.html", base))
	temps["forgot"] = template.Must(template.ParseFiles(dir+"forgot.html", base))
	temps["recover"] = template.Must(template.ParseFiles(dir+"recover.html", base))
	temps["small"] = template.Must(template.ParseFiles(dir+"small.html", base))

	log.Printf("Templates loaded\n")
}

func init() {
	loadTemplates()
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
		fmt.Fprintf(wr, "ERROR\n")
	}
}
