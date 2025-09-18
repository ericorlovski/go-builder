package model

type Field struct {
	Name string
	Type string
}

type StructMeta struct {
	PackageName string
	Name        string
	Fields      []Field
}
