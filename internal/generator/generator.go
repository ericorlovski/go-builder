package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ericorlovski/go-builder/internal/model"
)

const builderTpl = `
package {{.PackageName}}

type {{.Name}}Builder struct {
	{{- range .Fields}}
	{{lower .Name}} {{.Type}}
	{{- end}}
}

func New{{.Name}}Builder() *{{.Name}}Builder {
	return &{{.Name}}Builder{
		{{- range .Fields}}
		{{- if .Default}}
		{{lower .Name}}: {{renderDefault .Type .Default}},
		{{- end}}
		{{- end}}
	}
}

{{range .Fields}}
func (b *{{$.Name}}Builder) With{{.Name}}(v {{.Type}}) *{{$.Name}}Builder {
	b.{{lower .Name}} = v
	return b
}
{{end}}

func (b *{{.Name}}Builder) Build() {{.Name}} {
	// Required checks
	{{- range .Fields}}
		{{- if .Required}}
	if {{zeroCheck .Type (printf "b.%s" (lower .Name))}} {
		panic("{{$.Name}}.{{.Name}} is required")
	}
		{{- end}}
	{{- end}}

	obj := {{.Name}}{}

	// Assign fields
	{{- range .Fields}}
		{{- if .Omitempty}}
	if !({{zeroCheck .Type (printf "b.%s" (lower .Name))}}) {
		obj.{{.Name}} = b.{{lower .Name}}
	}
		{{- else}}
	obj.{{.Name}} = b.{{lower .Name}}
		{{- end}}
	{{- end}}

	return obj
}
`

func Generate(meta *model.StructMeta) (string, error) {
	funcs := template.FuncMap{
		"lower": func(s string) string {
			if len(s) == 0 {
				return s
			}
			if strings.ToUpper(s) == s {
				// for cases like ID, UUID
				return strings.ToLower(s)
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
		"renderDefault": func(typ string, def string) string {
			switch typ {
			case "int", "int64", "float64":
				return def
			case "bool":
				if def == "true" || def == "false" {
					return def
				}
				return "false"
			default:
				return fmt.Sprintf("%q", def)
			}
		},
		"zeroCheck": func(typ, varName string) string {
			switch typ {
			case "int", "int64", "float64":
				return fmt.Sprintf("%s == 0", varName)
			case "string":
				return fmt.Sprintf("%s == \"\"", varName)
			case "bool":
				return fmt.Sprintf("!%s", varName)
			default:
				return fmt.Sprintf("%s == nil", varName)
			}
		},
	}

	tmpl, err := template.New("builder").Funcs(funcs).Parse(builderTpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, meta)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
