package values

// Refers to an element inside of a map
type Map struct {
	valueType
	mapValue Any
	key      Any
}

// THe map that contains the value
func (m Map) Map() Any {
	return m.mapValue
}

// The key of the element that is being referred to
func (m Map) Key() Any {
	return m.key
}

func NewMap(mapValue, key Any) Map {
	return Map{
		mapValue: mapValue,
		key:      key,
	}
}
