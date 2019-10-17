package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("NASA WorldWind demo")

	// root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, path is: %q\n", html.EscapeString(r.URL.Path))
	})

	// static content handler
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// start serving
	log.Fatal(http.ListenAndServeTLS(":8080", "/h/certbot/rasp.beyondbroadcast.com/fullchain.pem",
		"/h/certbot/rasp.beyondbroadcast.com/privkey.pem", nil))
}
