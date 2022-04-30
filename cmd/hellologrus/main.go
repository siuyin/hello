package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetLevel(log.TraceLevel)
	log.Println("Gerbau")
	log.WithFields(log.Fields{"mod": "main"}).Info("log R us")

	params := []string{"brown", "fox", "lazy", "dog"}
	log.WithFields(log.Fields{"mod": "gerbau", "param": params}).Warn("log R us")

	name := struct {
		Name string
		Age  int
	}{"Siu Yin", 59}
	log.WithFields(log.Fields{"mod": "name", "name": name}).Info("my name")
	log.WithFields(log.Fields{"mod": "terpau"}).Fatal("terp crashed")
}
