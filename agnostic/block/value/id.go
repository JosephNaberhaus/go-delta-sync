package value

// A value within a name/id/variable
type Id struct {
	valueType
	methodIndependent
	name string
}

func (v Id) Name() string {
	return v.name
}

func NewId(name string) Id {
	return Id{name: name}
}
