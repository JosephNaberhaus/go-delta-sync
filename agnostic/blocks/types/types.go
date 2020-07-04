package types

type Any interface {
	isTypeType()
}

// Embed in any struct to mark it as a type of type
type typeType struct{}

func (t typeType) isTypeType() {}
