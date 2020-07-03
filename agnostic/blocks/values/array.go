package values

// Refers to an element contained by an array
type Array struct {
	valueType
	array Any
	index Any
}

func (a Array) Array() Any {
	return a.array
}

func (a Array) Index() Any {
	return a.index
}

func NewArray(array, index Any) Array {
	return Array{
		array: array,
		index: index,
	}
}
