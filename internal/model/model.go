package model

type Field struct {
	Name    string
	Type    string
	Default *string // nil if no default
}

type StructMeta struct {
	PackageName string
	Name        string
	Fields      []Field
}
