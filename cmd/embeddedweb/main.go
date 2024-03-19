// This example embeds a filesystem with the go binary.
// See pkg.go.dev/embed for detailed documentation.

package main

import (
	"log"
	"net/http"

	"github.com/siuyin/hello/cmd/embeddedweb/public"
)

func main() {
	http.Handle("/", http.FileServer(http.FS(public.Content))) // import "github.com/siuyin/hello/cmd/embeddedweb/public"
	//http.Handle("/", http.FileServer(http.Dir("public"))) // uncomment for quick develompent
	log.Fatal(http.ListenAndServe(":8080", nil))
}
