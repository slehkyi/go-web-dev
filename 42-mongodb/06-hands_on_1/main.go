package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/slehkyi/go-web-dev/42-mongodb/06-hands_on_1/controllers"
	"github.com/slehkyi/go-web-dev/42-mongodb/06-hands_on_1/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	s := map[string]models.User{}
	return s
}
