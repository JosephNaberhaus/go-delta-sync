package values

// Refers to an element contained by an array
type ArrayStruct struct {
	array Value
	index Value
}

func (a ArrayStruct) Array() Value {
	return a.array
}

func (a ArrayStruct) Index() Value {
	return a.index
}

func Array(array, index Value) ArrayStruct {
	return ArrayStruct{
		array: array,
		index: index,
	}
}
