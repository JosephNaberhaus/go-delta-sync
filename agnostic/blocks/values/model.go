package values

// Refers to a property that is part of another model
type ModelField struct {
	modelName string
	field     Value
}

// The name of the variable/reference that contains the model
func (m ModelField) ModelName() string {
	return m.modelName
}

func (m ModelField) Field() Value {
	return m.field
}

func NewModelField(modelName string, field Value) ModelField {
	return ModelField{
		modelName: modelName,
		field:     field,
	}
}
