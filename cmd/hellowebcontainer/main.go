package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	runWebServer()
}

// runWebServer is a function that uses net/http to run a webserver on port 8080
func runWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})
	http.ListenAndServe(":8080", nil)
}
