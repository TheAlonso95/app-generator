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

	// Generate Supabase Edge Function
	outputFile := "supabase_function.ts"
	port := 8000

	err = codegen.GenerateDenoFunction(swagger, outputFile, port)
	if err != nil {
		log.Fatalf("Error generating code: %v", err)
	}

	fmt.Println("🎉 Supabase Edge Function with dynamic routes generated successfully!")
}
