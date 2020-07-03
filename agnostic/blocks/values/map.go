package values

// Refers to an element inside of a map
type Map struct {
	mapValue Value
	key      Value
}

// THe map that contains the value
func (m Map) Map() Value {
	return m.mapValue
}

// The key of the element that is being referred to
func (m Map) Key() Value {
	return m.key
}

func NewMap(mapValue, key Value) Map {
	return Map{
		mapValue: mapValue,
		key:      key,
	}
}
