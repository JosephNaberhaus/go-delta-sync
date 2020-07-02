package types

type ArrayTypeDescription struct {
	elementDescription TypeDescription
}

func NewArrayTypeDescription(elementDescription TypeDescription) ArrayTypeDescription {
	return ArrayTypeDescription{elementDescription: elementDescription}
}

func (a ArrayTypeDescription) Classification() TypeClassification {
	return ArrayClassification
}

func (a ArrayTypeDescription) Value() string {
	return "[]" + a.Value()
}

func (a *ArrayTypeDescription) ElementDescription() TypeDescription {
	return a.elementDescription
}