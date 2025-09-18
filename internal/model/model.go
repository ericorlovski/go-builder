package model

type Field struct {
	Name      string
	Type      string
	Default   *string
	Required  bool
	Omitempty bool
	Validate  *string
}

type StructMeta struct {
	PackageName string
	Name        string
	Fields      []Field
}
