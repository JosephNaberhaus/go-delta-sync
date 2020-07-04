package types

// Represents a map from keys to values. A map should be able to map from any
// key to any value. However, language implementations are allowed to have
// limitations on when two keys are considered equal. However, the limitation
// must be mentioned on the implementations readme
type Map struct {
	typeType
	key, value Any
}

func NewMapTypeDescription(keyDescription, valueDescription Any) Map {
	return Map{key: keyDescription, value: valueDescription}
}

func (m Map) KeyDescription() Any {
	return m.key
}

func (m Map) ValueDescription() Any {
	return m.key
}
