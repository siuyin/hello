package main

import (
	"log"
	"os"

	_ "cldf"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	hostname := ""
	if localOnly := os.Getenv("LOCAL_ONLY"); localOnly == "true" {
		hostname = "127.0.0.1"
	}

	if err := funcframework.StartHostPort(hostname, port); err != nil {
		log.Fatal("funcFramework.StartHostPort: %v\n", err)
	}
}
