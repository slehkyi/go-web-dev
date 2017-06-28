package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	j := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corola",
		Doors:        4,
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}

	sages := []sage{b, j, g}
	cars := []car{c, f}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
