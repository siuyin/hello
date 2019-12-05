package main

import (
	"log"
	"net/http"
)

func main() {
	// serves content under html
	log.Println("secure web starting.")
	http.Handle("/", http.FileServer(http.Dir("html")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
