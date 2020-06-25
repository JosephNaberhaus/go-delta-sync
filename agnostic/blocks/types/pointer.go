package types

type PointerTypeDescription struct {
	valueTypeDescription TypeDescription // The type description of the value that is pointed to
}

func (p PointerTypeDescription) Classification() TypeClassification {
	return PointerClassification
}

func (p PointerTypeDescription) Value() string {
	return "*" + p.valueTypeDescription.Value()
}

func (p PointerTypeDescription) ValueTypeDescription() TypeDescription {
	return p.valueTypeDescription
}