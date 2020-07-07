package value

// Refers to a field that is a part of the model whose method is being called
type OwnField struct {
	valueType
	methodDependent
	field Any
}

func (o OwnField) Field() Any {
	return o.field
}

func NewOwnField(field Any) OwnField {
	return OwnField{field: field}
}
