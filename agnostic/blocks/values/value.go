package values

// Anything that represents a value that exists within the context of a
// methods body
type Any interface {
	isValueType()
}

// Embed in any struct to mark it as a type of value
type valueType struct{}

func (v valueType) isValueType() {}
