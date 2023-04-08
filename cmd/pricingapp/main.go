package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siuyin/dflt"
	"github.com/siuyin/hello/opa-pricing-app/db"
)

func main() {
	db := db.Init(dflt.EnvString("DB_NAME", "Pricing"))
	defer db.Close()

	fmt.Println("initialised")
	start := time.Now()
	buf := bytes.NewBufferString(`SKU,Description,Price
a,"apples,qty:5",1.99
b,"boy's pants,sm",3.45
a123,"almonds,20g",2.34`)
	db.Load(buf)

	buf = bytes.NewBufferString(`SKU,Description,Price
aa,"alpha milk,1L",2.99
bb,"boo baby formula,1kg",12.99`)
	db.Load(buf)

	val, err := db.Get("a")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(val.Value()))

	db.Dump(os.Stdout)
	fmt.Printf("run duration: %f seconds", time.Now().Sub(start).Seconds())
	//select {}
}
