package main

import (
	"text/template"
	"os"
	"log"
)

type person struct {
	Name string
	Age int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := doubleZero{
		person{
			Name: "James Bond",
			Age: 42,
		},
		false,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
