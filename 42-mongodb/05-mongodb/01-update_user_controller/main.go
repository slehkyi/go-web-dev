package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/slehkyi/go-web-dev/42-mongodb/05-mongodb/01-update_user_controller/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

