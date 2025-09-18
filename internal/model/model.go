package model

type Field struct {
	Name string
	Type string
}

type StructMeta struct {
	Name   string
	Fields []Field
}
