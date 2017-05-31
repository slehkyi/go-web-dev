package temp

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password string
	First string
	Last string
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
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
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
		u := user{un, p, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}