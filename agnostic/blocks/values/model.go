package values

// Refers to a property that is part of another model
type ModelFieldStruct struct {
	modelName string
	field     Value
}

// The name of the variable/reference that contains the model
func (m ModelFieldStruct) ModelName() string {
	return m.modelName
}

func (m ModelFieldStruct) Field() Value {
	return m.field
}

func ModelField(modelName string, field Value) ModelFieldStruct {
	return ModelFieldStruct{
		modelName: modelName,
		field:     field,
	}
}
