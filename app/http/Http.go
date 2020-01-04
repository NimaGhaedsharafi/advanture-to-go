package main

import (
	"fmt"
	"net/http"
	"testing/contact"
	"testing/template"
)

var (
	tmpl = template.Setup()
)

func Serve(adrs string) {
	http.HandleFunc("/", index)
	e := http.ListenAndServe(adrs, nil)
	if e != nil {
		panic("Couldn't start the server Error: " + e.Error())
	}
	fmt.Printf("It listens to %s\n", adrs)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*contact.Contact }{contact.GetAll()})
}
