package main

import (
	"fmt"
	"log"

	"github.com/generator/internal/fileparser"
)

func main() {
	swagger, err := fileparser.ReadYAML("swagger.yaml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Swagger API: %s\nTitle: %s\nVersion: %s\n", swagger.OpenAPI, swagger.Info.Title, swagger.Info.Version)
}
