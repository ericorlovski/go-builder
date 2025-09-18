package model

type Field struct {
	Name      string
	Type      string
	Default   *string
	Required  bool
	Omitempty bool
}

type StructMeta struct {
	PackageName string
	Name        string
	Fields      []Field
}
