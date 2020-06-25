package types

type ArrayTypeDescription struct {
	elementDescription TypeDescription
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