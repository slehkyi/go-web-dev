package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/slehkyi/go-web-dev/42-mongodb/04-mvc/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

