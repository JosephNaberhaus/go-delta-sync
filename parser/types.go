package parser

type Type struct {
	PackagePath, Name string
}

type Field struct {
	Name, Type string
}

type Struct struct {
	Name   string
	Fields []Field
}
