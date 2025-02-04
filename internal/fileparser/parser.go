package fileparser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// APIMethod represents a single API method inside a route
type APIMethod struct {
	Summary   string               `yaml:"summary"`
	Requests  map[string]SchemaRef `yaml:"requestBody,omitempty"`
	Responses map[string]Response  `yaml:"responses"`
}

// SchemaRef represents a reference to a schema definition
type SchemaRef struct {
	Content map[string]MediaType `yaml:"content"`
}

// MediaType holds the schema for a request or response
type MediaType struct {
	Schema Schema `yaml:"schema"`
}

// Schema represents the actual structure of a request or response
type Schema struct {
	Type       string            `yaml:"type"`
	Properties map[string]Schema `yaml:"properties,omitempty"`
}

// Response represents an API response structure
type Response struct {
	Description string               `yaml:"description"`
	Content     map[string]MediaType `yaml:"content,omitempty"`
}

// APIRoutes represents all API paths with their corresponding methods
type APIRoutes map[string]map[string]APIMethod

// Swagger represents a Swagger file structure
type Swagger struct {
	OpenAPI string    `yaml:"openapi"`
	Info    APIInfo   `yaml:"info"`
	Paths   APIRoutes `yaml:"paths"`
}

// APIInfo stores API title and version information
type APIInfo struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

// ReadYAML reads a Swagger YAML file and parses it
func ReadYAML(filePath string) (*Swagger, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var swagger Swagger
	if err := yaml.Unmarshal(data, &swagger); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &swagger, nil
}
