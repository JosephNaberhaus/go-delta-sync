package values

// A literal or reference to a value that exists inside of a methods body
type Value interface {
	valueType()
}

// Embed in any struct to mark it as a type of value
type valueType struct{}

func (v valueType) valueType() {}
