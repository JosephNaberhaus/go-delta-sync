package value

type Any interface {
	isValueType()
	IsMethodDependent() bool // True if the value has to be in the context of a method
}

type isValueType struct{}

func (v isValueType) isValueType() {}

type isMethodDependent struct{}

func (m isMethodDependent) IsMethodDependent() bool {
	return true
}

type isMethodIndependent struct{}

func (m isMethodIndependent) IsMethodDependent() bool {
	return false
}
