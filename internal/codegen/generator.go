package codegen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/generator/internal/fileparser"
)

// ToLower converts strings to lowercase (needed for Deno Oak methods)
func ToLower(s string) string {
	return strings.ToLower(s)
}

// GenerateResponseExample creates a sample JSON response from Swagger definitions
func GenerateResponseExample(responses map[string]fileparser.Response) string {
	for _, response := range responses {
		// Check if response content exists
		if len(response.Content) > 0 {
			for _, mediaType := range response.Content {
				if mediaType.Schema.Type == "object" {
					return generateObjectExample(mediaType.Schema.Properties)
				}
			}
		}
	}
	return `{}`
}

// generateObjectExample converts a Swagger object schema into a TypeScript JSON response
func generateObjectExample(properties map[string]fileparser.Schema) string {
	if len(properties) == 0 {
		return "{}" // Return empty object if no properties exist
	}

	result := "{\n"
	for key, prop := range properties {
		result += fmt.Sprintf(`  "%s": %s,`, key, getDefaultValue(prop.Type))
	}
	result = strings.TrimSuffix(result, ",") + "\n}" // Remove trailing comma
	return result
}

// getDefaultValue returns a sample JSON value based on its type
func getDefaultValue(schemaType string) string {
	switch schemaType {
	case "string":
		return `"example_string"` // More descriptive placeholder
	case "integer":
		return "0" // Use a realistic default instead of 42
	case "boolean":
		return "false" // Explicitly false by default
	case "array":
		return "[]" // Represents an empty array
	case "object":
		return "{}" // Represents an empty object
	default:
		return "null" // Default for unknown types
	}
}

// FunctionData represents the template data for a generated function
type FunctionData struct {
	Port   int
	Routes fileparser.APIRoutes
}

// ExtractBaseRoute gets the first segment of a route
func ExtractBaseRoute(route string) string {
	segments := strings.Split(strings.Trim(route, "/"), "/")
	if len(segments) > 0 {
		return segments[0]
	}
	return "default"
}

// GenerateDenoFunctions reads a template file and generates Supabase Edge Functions
func GenerateDenoFunctions(swagger *fileparser.Swagger, functionsDir string, port int) error {
	// Read the template file
	templatePath := "templates/deno_function.tmpl"
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	funcMap := template.FuncMap{
		"ToLower":                 ToLower,
		"GenerateResponseExample": GenerateResponseExample,
	}

	tmpl, err := template.New("deno").Funcs(funcMap).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Ensure `functions/` directory exists
	if err := os.MkdirAll(functionsDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create functions directory: %w", err)
	}

	// Track generated functions to avoid duplicate folders
	generatedFolders := make(map[string]bool)

	fmt.Println("\n🚦 Swagger Paths:")
	for route := range swagger.Paths {
		fmt.Printf("  - %s\n", route)
	}
	fmt.Println()

	for route, methods := range swagger.Paths {
		baseRoute := ExtractBaseRoute(route)
		functionFolder := filepath.Join(functionsDir, baseRoute)
		functionFile := filepath.Join(functionFolder, "index.ts")

		// Create function folder if it doesn't exist
		if !generatedFolders[baseRoute] {
			if err := os.MkdirAll(functionFolder, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create function folder: %w", err)
			}
			generatedFolders[baseRoute] = true
		}

		// Prepare data for the template
		data := FunctionData{
			Port:   port,
			Routes: fileparser.APIRoutes{route: methods},
		}

		fmt.Printf("\n📝 Function Data:\n")
		fmt.Printf("  Port: %d\n", data.Port)
		fmt.Printf("  Routes:\n")
		for r, m := range data.Routes {
			fmt.Printf("    %s: %+v\n", r, m)
		}
		fmt.Println()

		// Generate the function file
		file, err := os.Create(functionFile)
		if err != nil {
			return fmt.Errorf("failed to create function file: %w", err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, data); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}

		fmt.Println("✅ Generated Supabase Edge Function:", functionFile)
	}

	return nil
}
