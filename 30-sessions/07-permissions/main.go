package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First string
	Last string
	Role string
}

var tpl *template.Template
var dbUsers = map[string]user{} // userID, user{}
var dbSessions = map[string]string{} // sessionID, userID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	//process form submission
	if r.Method == http.MethodPost {
		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		role := r.FormValue("role")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already exists", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l, role}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "User doesn't exist", http.StatusForbidden)
			return
		}
		// does the entered password match the stored one?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Password does not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	// delete session
	delete(dbSessions, c.Value)
	// remove cookie
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}