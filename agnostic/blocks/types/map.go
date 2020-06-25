package types

type MapTypeDescription struct {
	keyDescription TypeDescription
	valueDescription TypeDescription
}

func (m MapTypeDescription) Classification() TypeClassification {
	return MapClassification
}

func (m MapTypeDescription) Value() string {
	return "map[" + m.keyDescription.Value() + "]" + m.valueDescription.Value()
}

func (m MapTypeDescription) KeyDescription() TypeDescription {
	return m.keyDescription
}

func (m MapTypeDescription) ValueDescription() TypeDescription {
	return m.valueDescription
}