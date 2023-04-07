package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/siuyin/hello/opa-pricing-app/db"
)

func main() {
	db := db.Init("pricing")
	defer db.Close()

	fmt.Println("initialised")
	start := time.Now()
	buf := bytes.NewBufferString(`SKU,Description,Price
a,"apples,qty:5",1.99
b,"boy's pants,sm",3.45
a123,"almonds,20g",2.34`)
	db.Load(buf)

	buf = bytes.NewBufferString(`SKU,Description,Price
aa,"alpha milk,1L",2.99`)
	db.Load(buf)

	db.Dump(os.Stdout)
	fmt.Printf("run duration: %f seconds", time.Now().Sub(start).Seconds())
}
