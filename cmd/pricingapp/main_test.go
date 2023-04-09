package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/siuyin/dflt"
	"github.com/siuyin/hello/opa-pricing-app/db"
)

func Example() {
	db := db.Init(dflt.EnvString("DB_NAME", "Pricing"))
	defer db.Close()

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

	// Output:
	// {"SKU":"a","Description":"apples,qty:5","Price":1.99}
	// SKU,Description,Price
	// a,"apples,qty:5",1.99
	// b,"boy's pants,sm",3.45
	// a123,"almonds,20g",2.34
	// aa,"alpha milk,1L",2.99
	// bb,"boo baby formula,1kg",12.99
}
