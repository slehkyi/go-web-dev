package config

import (
	"gopkg.in/mgo.v2"
	"fmt"
	_ "github.com/lib/pq"
)

//database
var DB *mgo.Database

// collections
var Books *mgo.Collection

func init() {
	var err error
	s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")
	fmt.Println("You connected to your mongo database.")
}
