package types

// Represents a map from keys to values. A map should be able to map from any
// key to any value. However, language implementations are allowed to have
// limitations on when two keys are considered equal. However, the limitation
// must be mentioned on the implementations readme
type Map struct {
	typeType
	key, value Any
}

func NewMap(key, value Any) Map {
	return Map{key: key, value: value}
}

func (m Map) Key() Any {
	return m.key
}

func (m Map) Value() Any {
	return m.key
}
