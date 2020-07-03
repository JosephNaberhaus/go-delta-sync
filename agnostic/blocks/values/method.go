package values

// Refers to a field that is a part of the model whose method is being called
type OwnField struct {
	valueType
	field Any
}

func (o OwnField) Field() Any {
	return o.field
}

func NewOwnField(field Any) OwnField {
	return OwnField{field: field}
}
