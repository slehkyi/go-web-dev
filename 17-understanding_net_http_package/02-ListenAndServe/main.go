package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Any code you want in this func")
}

func main()  {
	var h hotdog
	http.ListenAndServe(":8080", h)
}