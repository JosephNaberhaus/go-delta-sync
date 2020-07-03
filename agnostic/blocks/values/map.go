package values

// Refers to an element inside of a map
type MapStruct struct {
	mapValue Value
	key      Value
}

// THe map that contains the value
func (m MapStruct) Map() Value {
	return m.mapValue
}

// The key of the element that is being referred to
func (m MapStruct) Key() Value {
	return m.key
}

func Map(mapValue, key Value) MapStruct {
	return MapStruct{
		mapValue: mapValue,
		key:      key,
	}
}
