package value

// Refers to an element contained by an array
type ArrayElement struct {
	valueType
	array Any
	index Any
}

func (a ArrayElement) Array() Any {
	return a.array
}

func (a ArrayElement) Index() Any {
	return a.index
}

func NewArrayElement(array, index Any) ArrayElement {
	return ArrayElement{
		array: array,
		index: index,
	}
}
