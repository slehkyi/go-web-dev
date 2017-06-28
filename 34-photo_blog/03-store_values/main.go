package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

type UserPics struct {
	sID     string
	PicName string
}

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
	c = appendString(w, c)
	xs := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
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

func appendString(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	//picture names
	pic1 := "pic1.jpg"
	pic2 := "pic2.jpg"
	pic3 := "pic3.jpg"

	//add pics to cookie's value if don't exist
	s := c.Value
	if !strings.Contains(s, pic1) {
		s += "|" + pic1
	}
	if !strings.Contains(s, pic2) {
		s += "|" + pic2
	}
	if !strings.Contains(s, pic3) {
		s += "|" + pic3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
