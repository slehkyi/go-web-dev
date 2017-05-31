package main

import (
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at foo: ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar:", r.Method)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred:", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}