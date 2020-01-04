package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	. "testing/contact"
	"testing/template"
)

var (
	listenAdrs = flag.String("adrs", os.Getenv("HTTP_ADRS"), "HTTP address to listen on")
	tmpl       = template.Setup()
)

func main() {
	flag.Parse()

	http.HandleFunc("/", index)
	e := http.ListenAndServe(*listenAdrs, nil)
	if e != nil {
		panic("Couldn't start the server")
	}
	fmt.Printf("It listens to %s\n", *listenAdrs)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*Contact }{GetAll()})
}
