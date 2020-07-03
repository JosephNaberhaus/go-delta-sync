package values

// Refers to a field that is a part of the model whose method is being called
type OwnFieldStruct struct {
	field Value
}

func (o OwnFieldStruct) Field() Value {
	return o.field
}

func OwnField(field Value) OwnFieldStruct {
	return OwnFieldStruct{field: field}
}
