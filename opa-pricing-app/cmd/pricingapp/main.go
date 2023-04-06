package main

import (
	"fmt"
	"pricing/db"
)

func main() {
	db := db.Init("pricing")
	defer db.Close()

	fmt.Println("initialised")
	select {}
}
