package value

type Any interface {
	isValueType()
	IsMethodDependent() bool // True if the value has to be in the context of a method
}

type valueType struct{}

func (v valueType) isValueType() {}

type methodDependent struct{}

func (m methodDependent) IsMethodDependent() bool {
	return true
}

type methodIndependent struct{}

func (m methodIndependent) IsMethodDependent() bool {
	return false
}
