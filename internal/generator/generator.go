package generator

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/username/go-builder/internal/model"
)

const builderTpl = `
type {{.Name}}Builder struct {
	{{- range .Fields}}
	{{lower .Name}} {{.Type}}
	{{- end}}
}

func New{{.Name}}Builder() *{{.Name}}Builder {
	return &{{.Name}}Builder{}
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
			return strings.ToLower(s[:1]) + s[1:]
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
