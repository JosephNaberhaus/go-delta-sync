package value

type IntToString struct {
	isValueType
	intValue Any
}

func (i IntToString) IntValue() Any {
	return i.intValue
}

func (i IntToString) IsMethodDependent() bool {
	return i.intValue.IsMethodDependent()
}

func NewIntToString(intValue Any) IntToString {
	return IntToString{
		intValue: intValue,
	}
}
