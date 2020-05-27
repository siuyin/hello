package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("HTML Interaction")
	http.HandleFunc("/", rootHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./tmpl/*/*.html"))
	t.Execute(w, nil)
}
