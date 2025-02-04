package main

import (
	"fmt"
	"log"

	"github.com/generator/internal/codegen"
	"github.com/generator/internal/fileparser"
)

func main() {
	// Read Swagger YAML
	swagger, err := fileparser.ReadYAML("swagger.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML: %v", err)
	}

	// Define Supabase functions folder
	functionsDir := "functions"
	port := 8000

	// Generate Edge Functions
	err = codegen.GenerateDenoFunctions(swagger, functionsDir, port)
	if err != nil {
		log.Fatalf("Error generating code: %v", err)
	}

	fmt.Println("🎉 Supabase Edge Functions successfully generated!")
}
