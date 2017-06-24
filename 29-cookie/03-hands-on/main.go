package main

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
	"io"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:"my-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)

}
