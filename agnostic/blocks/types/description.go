package types

type TypeClassification int

const (
	ArrayClassification   TypeClassification = iota // A types representing a array of elements
	MapClassification                               // A types representing a collection of key-value pairs
	PointerClassification                           // A types representing a pointer to another value
	BaseClassification                              // A types representing one of the built-in base types of Go
	IdClassification                                // A types representing the name of another model
)

type TypeDescription interface {
	Classification() TypeClassification
	Value() string // Equivalent to the Go representation of the types
}
