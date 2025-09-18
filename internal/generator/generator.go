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
	return {{.Name}}{
		{{- range .Fields}}
		{{.Name}}: b.{{lower .Name}},
		{{- end}}
	}
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
