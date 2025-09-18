package parser

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/ericorlovski/go-builder/internal/model"
)

func ParseStruct(filename, structName string) (*model.StructMeta, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var meta model.StructMeta
	meta.Name = structName
	meta.PackageName = node.Name.Name

	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok || ts.Name.Name != structName {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return false
		}

		for _, field := range st.Fields.List {
			if len(field.Names) == 0 {
				continue
			}
			name := field.Names[0].Name
			typ := exprToString(field.Type)
			meta.Fields = append(meta.Fields, model.Field{Name: name, Type: typ})
		}
		return false
	})

	return &meta, nil
}

func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	default:
		return "interface{}"
	}
}
