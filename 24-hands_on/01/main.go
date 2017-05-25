package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}
