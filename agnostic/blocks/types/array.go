package types

type Array struct {
	elementDescription TypeDescription
}

func NewArrayDescription(elementDescription TypeDescription) ArrayTypeDescription {
	return ArrayTypeDescription{elementDescription: elementDescription}
}

func (a ArrayTypeDescription) Classification() TypeClassification {
	return ArrayClassification
}

func (a ArrayTypeDescription) Value() string {
	return "[]" + a.elementDescription.Value()
}

func (a *ArrayTypeDescription) ElementDescription() TypeDescription {
	return a.elementDescription
}
