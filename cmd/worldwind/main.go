package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
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
	certPath := dflt.EnvString("CERT_PATH", "/h/certbot/rasp.beyondbroadcast.com/fullchain.pem")
	keyPath := dflt.EnvString("KEY_PATH", "/h/certbot/rasp.beyondbroadcast.com/privkey.pem")
	port := dflt.EnvString("PORT", ":8080")
	log.Fatal(http.ListenAndServeTLS(port, certPath, keyPath, nil))
}
