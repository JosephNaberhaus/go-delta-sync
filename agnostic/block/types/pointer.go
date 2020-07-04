package types

// Represents a pointer type. Languages that don't support pointers should
// implement this type in some way that mimics pointers
type Pointer struct {
	typeType
	value Any // The type of the value that is pointed to
}

func NewPointer(value Any) Pointer {
	return Pointer{value: value}
}

func (p Pointer) Value() Any {
	return p.value
}
