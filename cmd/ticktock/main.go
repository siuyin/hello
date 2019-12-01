package main

import (
	"fmt"
	"time"

	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "dev")
	fmt.Println("ticktock prints the time once a second")
	for {
		fmt.Printf("%s: %s\n", stage, time.Now().Format("2006-01-02 15:04:05.00 MST"))
		time.Sleep(time.Second)
	}
}
