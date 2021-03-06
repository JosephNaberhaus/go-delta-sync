package value

// Refers to a property that is part of another model
type ModelField struct {
	isValueType
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

func (m ModelField) IsMethodDependent() bool {
	return m.field.IsMethodDependent()
}

func NewModelField(modelName string, field Any) ModelField {
	return ModelField{
		modelName: modelName,
		field:     field,
	}
}
