package main

import (
	"fmt"
	"time"

	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "dev")
	msg := dflt.EnvString("MSG", "dev-msg")
	fmt.Printf("%s ticktock prints the time once a second", stage)
	for {
		fmt.Printf("%s: %s: %s\n", stage, msg, time.Now().Format("2006-01-02 15:04:05.00 MST"))
		time.Sleep(time.Second)
	}
}
