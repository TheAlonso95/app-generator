package fileparser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Swagger represents a basic structure of a Swagger file (extend as needed)
type Swagger struct {
	OpenAPI string `yaml:"openapi"`
	Info    struct {
		Title   string `yaml:"title"`
		Version string `yaml:"version"`
	} `yaml:"info"`
}

// ReadYAML reads a Swagger YAML file and parses it into a Swagger struct
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
