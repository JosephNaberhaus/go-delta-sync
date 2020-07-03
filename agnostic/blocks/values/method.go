package values

// Refers to a field that is a part of the model whose method is being called
type OwnField struct {
	valueType
	field Value
}

func (o OwnField) Field() Value {
	return o.field
}

func NewOwnField(field Value) OwnField {
	return OwnField{field: field}
}
