package template

import (
	"html/template"
	"log"
	"path/filepath"
	"testing/util"
)

func Setup() *template.Template {
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{"Strong": util.Strong})
	_, e := tmpl.ParseGlob(filepath.Join(".", "*.html"))
	if e != nil {
		log.Fatalln(e)
		panic("Template setup failed")
	}

	return tmpl
}
