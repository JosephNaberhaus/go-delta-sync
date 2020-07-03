package values

// Refers to a literal null/nil/empty value
type NullStruct struct{}

func Null() NullStruct {
	return NullStruct{}
}

// Refers to a literal string value
type StringStruct struct {
	value string
}

func (s StringStruct) Value() string {
	return s.value
}

func String(value string) StringStruct {
	return StringStruct{value: value}
}

// Refers a literal int value
type IntStruct struct {
	value int
}

func (i IntStruct) Value() int {
	return i.value
}

func Int(value int) IntStruct {
	return IntStruct{value: value}
}

// Refers to a literal floating point value
type FloatStruct struct {
	value float64
}

func (f FloatStruct) Value() float64 {
	return f.value
}

func Float(value float64) FloatStruct {
	return FloatStruct{value: value}
}

// Refers to a literal boolean value
type BoolStruct struct {
	value bool
}

func (b BoolStruct) Value() bool {
	return b.value
}

func Bool(value bool) BoolStruct {
	return BoolStruct{value: value}
}
