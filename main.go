package main

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"testing/connection"
	. "testing/contact"
	"testing/template"
)

var (
	listenAdrs = flag.String("adrs", os.Getenv("HTTP_ADRS"), "HTTP address to listen on")
	tmpl       = template.Setup()
	db         *sqlx.DB
)

func main() {
	flag.Parse()
	db := connection.Connection()

	if db == nil {
		os.Exit(2)
	}

	_, e := db.Exec("select * from contacts")

	if e != nil {
		os.Exit(3)
	}

	fmt.Printf("It listens to %s\n", *listenAdrs)
	http.HandleFunc("/", index)
	http.ListenAndServe(*listenAdrs, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", struct{ Contacts []*Contact }{GetAll()})
}
