package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // userID, user{}
var dbSessions = map[string]string{} // sessionID, userID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// get cookie, if cookie doesn't exist, create cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println(err)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
