package main

import (
	"net/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.StripPrefix("/pic/", http.FileServer(http.Dir(".")))))
}

