package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
