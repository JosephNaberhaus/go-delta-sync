package value

// Refers to a property that is part of another model
type ModelField struct {
	valueType
	modelName string
	field     Any
}

// The name of the variable/reference that contains the model
func (m ModelField) ModelName() string {
	return m.modelName
}

func (m ModelField) Field() Any {
	return m.field
}

func NewModelField(modelName string, field Any) ModelField {
	return ModelField{
		modelName: modelName,
		field:     field,
	}
}
