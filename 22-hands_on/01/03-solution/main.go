// Take the previous program and change it so that:
// func main uses http.Handle instead of http.HandleFunc
// Constraint: Do not change anything outside of func main

package main

import (
	"io"
	"net/http"
	"html/template"
	"log"
)

func main() {

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(myName))

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
