package types

// Represents the type of another model
type Model struct {
	typeType
	modelName string
}

func NewModel(modelName string) Model {
	return Model{modelName: modelName}
}

func (m Model) ModelName() string {
	return m.modelName
}
