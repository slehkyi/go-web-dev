package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.StripPrefix("/pic/", http.FileServer(http.Dir(".")))))
}
