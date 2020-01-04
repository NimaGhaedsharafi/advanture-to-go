package contact

import (
	"log"
	"testing/connection"
)

var (
	DB = connection.Connection()
)

type Contact struct {
	Id int
	Name string
	Tel string
	CreatedAt string `db:"created_at"`
}

func GetAll() (contacts []*Contact) {
	e := DB.Select(&contacts, "select * from contacts")

	if e != nil {
		log.Fatalln(e)
		return nil
	}

	return contacts
}
