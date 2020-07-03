package blocks

type Value interface{}

// Refers to a literal null/nil/empty value
type NullValueStruct struct{}

func NullValue() NullValueStruct {
	return NullValueStruct{}
}

// Refers to a literal string value
type StringValueStruct struct {
	value string
}

func (s StringValueStruct) Value() string {
	return s.value
}

func StringValue(value string) StringValueStruct {
	return StringValueStruct{value: value}
}

// Refers a literal int value
type IntValueStruct struct {
	value int
}

func (i IntValueStruct) Value() int {
	return i.value
}

func IntValue(value int) IntValueStruct {
	return IntValueStruct{value: value}
}

// Refers to a literal floating point value
type FloatValueStruct struct {
	value float64
}

func (f FloatValueStruct) Value() float64 {
	return f.value
}

func FloatValue(value float64) FloatValueStruct {
	return FloatValueStruct{value: value}
}

// Refers to a literal boolean value
type BoolValueStruct struct {
	value bool
}

func (b BoolValueStruct) Value() bool {
	return b.value
}

func BoolValue(value bool) BoolValueStruct {
	return BoolValueStruct{value: value}
}

// Refers to a property that is a part of the model whose method is being called
type OwnPropertyStruct struct {
	name string
}

func (o OwnPropertyStruct) Name() string {
	return o.name
}

func OwnProperty(name string) OwnPropertyStruct {
	return OwnPropertyStruct{name: name}
}

// Refers to a variable that is within scope of the current block
type VariableStruct struct {
	name string
}

func (v VariableStruct) Name() string {
	return v.name
}

func Variable(name string) VariableStruct {
	return VariableStruct{name: name}
}

// Refers to a property that is part of another model. Note that model name is the name of the variable/parameter that
// contains the model and not the name of the model itself
type ModelPropertyStruct struct {
	modelName, name string
}

func (m ModelPropertyStruct) ModelName() string {
	return m.modelName
}

func (m ModelPropertyStruct) Name() string {
	return m.name
}

func ModelProperty(modelName, name string) ModelPropertyStruct {
	return ModelPropertyStruct{
		modelName: modelName,
		name:      name,
	}
}

// Refers to an element contained by an array
type ArrayValueStruct struct {
	array Value
	index Value
}

func (a ArrayValueStruct) Array() Value {
	return a.array
}

func (a ArrayValueStruct) Index() Value {
	return a.index
}

func ArrayValue(array, index Value) ArrayValueStruct {
	return ArrayValueStruct{
		array: array,
		index: index,
	}
}

// Refers to an element inside of a map
type MapValueStruct struct {
	mapValue Value
	key      Value
}

func (m MapValueStruct) Map() Value {
	return m.mapValue
}

func (m MapValueStruct) Key() Value {
	return m.key
}

func MapValue(mapValue, key Value) MapValueStruct {
	return MapValueStruct{
		mapValue: mapValue,
		key:      key,
	}
}
