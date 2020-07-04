package types

// Represents a subset of base types in Go. These include only the types that
// are likely to be supported in other languages (with Javascript being a
// priority).
type Base int

func (b Base) isTypeType() {}

const (
	BaseInt Base = iota
	BaseInt32
	BaseInt64
	BaseFloat32
	BaseFloat64
	BaseBool
	BaseString
	NumberBaseTypes // IMPORTANT: Keep at end
)

// This constructor is only added so that it matches other types
func NewBase(base Base) Base {
	return base
}
