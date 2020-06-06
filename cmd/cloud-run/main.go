package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
)

func main() {
	port := dflt.EnvString("PORT", "8080")
	fmt.Println("service starting on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello\n")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
