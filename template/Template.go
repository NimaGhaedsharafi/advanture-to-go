package template

import (
	"html/template"
	"log"
	"path/filepath"
	"testing/Util"
)

func Setup() *template.Template {
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{"Strong": Util.Strong})
	_, e := tmpl.ParseGlob(filepath.Join(".", "*.html"))
	if e != nil {
		log.Fatalln(e)
		panic("Template setup failed")
	}

	return tmpl
}
