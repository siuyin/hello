package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
	"github.com/siuyin/hello/world"
)

func main() {
	subject := dflt.EnvString("SUBJECT", "Siu Yin")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s!\n", world.Greet(), subject)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
