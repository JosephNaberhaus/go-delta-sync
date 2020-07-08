package types

// Represents an array/slice type
type Array struct {
	typeType
	element Any
}

func NewArray(element Any) Array {
	return Array{element: element}
}

func (a Array) Element() Any {
	return a.element
}
