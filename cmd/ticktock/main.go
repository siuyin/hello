package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("ticktock prints the time once a second")
	for {
		fmt.Println(time.Now().Format("time: 2006-01-02 15:04:05.00 MST"))
		time.Sleep(time.Second)
	}
}
