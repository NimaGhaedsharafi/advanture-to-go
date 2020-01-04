package http

import (
	"net/http"
	"testing/contact"
	"testing/template"
)

var (
	tmpl = template.Setup()
)

func Serve(adrs *string) {
	http.HandleFunc("/", index)
	e := http.ListenAndServe(*adrs, nil)
	if e != nil {
		panic("Couldn't start the server Error: " + e.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*contact.Contact }{contact.GetAll()})
}
