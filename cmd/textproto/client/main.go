package main

import (
	"fmt"
	"log"
	"net/textproto"
	"time"

	"github.com/siuyin/dflt"
)

func main() {
	fmt.Printf("text protocol client demo\n")
	svrAddr := dflt.EnvString("SVR_ADDR", "localhost:3999")
	fmt.Printf("Attempting to connect to %s\n", svrAddr)
	c, err := textproto.Dial("tcp", svrAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	code, msg, err := c.ReadCodeLine(10) // expect a response of 10[0..9]
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %s\n", code, msg)

	time.Sleep(time.Second)
	id, err := c.Cmd("HELP")
	if err != nil {
		log.Fatal(err)
	}

	c.StartResponse(id)
	defer c.EndResponse(id)

	line, err := c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(line)
}
