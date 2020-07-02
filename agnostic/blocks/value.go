package blocks

type Value interface{}

// Refers to a literal string value
type StringValue struct {
	value string
}

func (s StringValue) Value() string {
	return s.value
}

// Refers a literal int value
type IntValue struct {
	value int
}

func (i IntValue) Value() int {
	return i.value
}

// Refers to a literal floating point value
type FloatValue struct {
	value float64
}

func (f FloatValue) Value() float64 {
	return f.value
}

// Refers to a property that is a part of the model whose method is being called
type OwnProperty struct {
	name string
}

func (o OwnProperty) Name() string {
	return o.name
}

// Refers to a variable that is within scope of the current block
type Variable = OwnProperty

// Refers to a property that is part of another model. Note that model name is the name of the variable/parameter that
// contains the model and not the name of the model itself
type ModelProperty struct {
	modelName, name string
}

func (m ModelProperty) ModelName() string {
	return m.modelName
}

func (m ModelProperty) Name() string {
	return m.name
}

// Refers to an element contained by an array
type ArrayValue struct {
	array Value
	index Value
}

func (a ArrayValue) Array() Value {
	return a.array
}

func (a ArrayValue) Index() Value {
	return a.index
}

func (m MapValue) Map() Value {
	return m.mapValue
}

// Refers to an element inside of a map
type MapValue struct {
	mapValue Value
	key      Value
}

func (m MapValue) Key() Value {
	return m.key
}
