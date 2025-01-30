package codegen

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/generator/internal/fileparser"
)

// Template for a Supabase Edge Function in Deno with Oak
const denoTemplate = `import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const router = new Router();
{{range $path, $methods := .Routes}}
{{range $method, $info := $methods}}
router.{{$method | ToLower}}("{{$path}}", (ctx) => {
  ctx.response.body = "{{.Summary}}";
});
{{end}}
{{end}}

const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());

console.log("Server running on http://localhost:{{.Port}}");
await app.listen({ port: {{.Port}} });
`

func ToLower(s string) string {
	return strings.ToLower(s)
}

// GenerateDenoFunction creates a Deno Oak server file based on Swagger
func GenerateDenoFunction(swagger *fileparser.Swagger, outputFile string, port int) error {
	funcMap := template.FuncMap{
		"ToLower": ToLower,
	}

	tmpl, err := template.New("deno").Funcs(funcMap).Parse(denoTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	data := struct {
		Title   string
		Version string
		Port    int
		Routes  map[string]map[string]struct {
			Summary string `yaml:"summary"`
		}
	}{
		Title:   swagger.Info.Title,
		Version: swagger.Info.Version,
		Port:    port,
		Routes:  swagger.Paths,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Println("✅ Generated Supabase Edge Function with dynamic routes:", outputFile)
	return nil
}
