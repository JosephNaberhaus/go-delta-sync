package values

// Refers to an element contained by an array
type Array struct {
	valueType
	array Value
	index Value
}

func (a Array) Array() Value {
	return a.array
}

func (a Array) Index() Value {
	return a.index
}

func NewArray(array, index Value) Array {
	return Array{
		array: array,
		index: index,
	}
}
