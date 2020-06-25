package types

type IdTypeDescription struct {
	id string
}

func (i IdTypeDescription) Classification() TypeClassification {
	return IdClassification
}

func (i IdTypeDescription) Value() string {
	return i.id
}

func (i IdTypeDescription) Id() string {
	return i.id
}