package main

import (
	"fmt"

	"github.com/siuyin/hello/opa-pricing-app/db"
)

func main() {
	db := db.Init("pricing")
	defer db.Close()

	fmt.Println("initialised")
	select {}
}
