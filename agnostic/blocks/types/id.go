package types

type IdTypeDescription struct {
	id string
}

func NewIdTypeDescription(id string) IdTypeDescription {
	return IdTypeDescription{id: id}
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
