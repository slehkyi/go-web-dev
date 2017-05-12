package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		42,
		13,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
