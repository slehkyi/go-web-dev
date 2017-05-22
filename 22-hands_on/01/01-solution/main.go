// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", myName)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is index")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Marley is a dog")
}

func myName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My name is Sergi")
}
