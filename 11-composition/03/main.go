package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Intro to Go", "4"},
				course{"CSCI-50", "Web with Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-60", "Advanced Go", "5"},
				course{"CSCI-70", "Advanced Web with Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
