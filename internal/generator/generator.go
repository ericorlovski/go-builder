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
// {{.Name}} sets the {{.Name}} field for {{$.Name}}
func (b *{{$.Name}}Builder) {{.Name}}(v {{.Type}}) *{{$.Name}}Builder {
	b.{{lower .Name}} = v
	return b
}
{{end}}

func (b *{{.Name}}Builder) Build() {{.Name}} {
	// Required checks
	{{- range .Fields}}
		{{- if .Required}}
	if {{isZero .Type (printf "b.%s" (lower .Name))}} {
		panic("{{$.Name}}.{{.Name}} is required")
	}
		{{- end}}
	{{- end}}

	// Validation checks
	{{- range .Fields}}
    	{{- if .Validate}}
    {{renderValidation (printf "b.%s" (lower .Name)) .Type .Validate}}
    	{{- end}}
	{{- end}}

	obj := {{.Name}}{}

	// Assign fields
	{{- range .Fields}}
		{{- if .Omitempty}}
	if {{notZero .Type (printf "b.%s" (lower .Name))}} {
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
		"isZero": func(typ, varName string) string {
			switch {
			case typ == "int" || typ == "int64" || typ == "float64":
				return fmt.Sprintf("%s == 0", varName)
			case typ == "string":
				return fmt.Sprintf("%s == \"\"", varName)
			case typ == "bool":
				return fmt.Sprintf("!%s", varName)
			case strings.HasPrefix(typ, "[]"):
				return fmt.Sprintf("len(%s) == 0", varName)
			case strings.HasPrefix(typ, "map["):
				return fmt.Sprintf("%s == nil || len(%s) == 0", varName, varName)
			case strings.HasPrefix(typ, "*"):
				return fmt.Sprintf("%s == nil", varName)
			default:
				return fmt.Sprintf("%s == nil", varName)
			}
		},
		"notZero": func(typ, varName string) string {
			switch {
			case typ == "int" || typ == "int64" || typ == "float64":
				return fmt.Sprintf("%s != 0", varName)
			case typ == "string":
				return fmt.Sprintf("%s != \"\"", varName)
			case typ == "bool":
				return fmt.Sprintf("%s", varName) // true остаётся
			case strings.HasPrefix(typ, "[]"):
				return fmt.Sprintf("len(%s) > 0", varName)
			case strings.HasPrefix(typ, "map["):
				return fmt.Sprintf("%s != nil && len(%s) > 0", varName, varName)
			case strings.HasPrefix(typ, "*"):
				return fmt.Sprintf("%s != nil", varName)
			default:
				return fmt.Sprintf("%s != nil", varName)
			}
		},
		"renderValidation": func(name, typ string, rule *string) string {
			if rule == nil {
				return ""
			}
			parts := strings.Split(*rule, ",")
			var checks []string
			for _, p := range parts {
				if strings.HasPrefix(p, "min=") {
					val := strings.TrimPrefix(p, "min=")
					checks = append(checks, fmt.Sprintf("%s < %s", name, val))
				}
				if strings.HasPrefix(p, "max=") {
					val := strings.TrimPrefix(p, "max=")
					checks = append(checks, fmt.Sprintf("%s > %s", name, val))
				}
				if p == "email" {
					checks = append(checks, fmt.Sprintf("!strings.Contains(%s, \"@\")", name))
				}
			}
			if len(checks) == 0 {
				return ""
			}
			return fmt.Sprintf("if %s { panic(\"validation failed for %s\") }", strings.Join(checks, " || "), name)
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
