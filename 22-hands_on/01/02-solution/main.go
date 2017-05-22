// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

package main

import (
	"io"
	"net/http"
	"html/template"
	"log"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", myName)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Marley is a dog")
}

func myName(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", "Sergi")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
