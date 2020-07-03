package values

// Anything that represents a value that exists within the context of a
// methods body
type Any interface {
	valueType()
}

// Embed in any struct to mark it as a type of value
type valueType struct{}

func (v valueType) valueType() {}
