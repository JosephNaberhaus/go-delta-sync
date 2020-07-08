package value

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
)

// Refers to a literal null/nil/empty value
type Null struct {
	isValueType
	isMethodIndependent
}

func NewNull() Null {
	return Null{}
}

// Refers to a literal string value
type String struct {
	isValueType
	isMethodIndependent
	value string
}

func (s String) Value() string {
	return s.value
}

func NewString(value string) String {
	return String{value: value}
}

// Refers a literal int value
type Int struct {
	isValueType
	isMethodIndependent
	value int
}

func (i Int) Value() int {
	return i.value
}

func NewInt(value int) Int {
	return Int{value: value}
}

// Refers to a literal floating point value
type Float struct {
	isValueType
	isMethodIndependent
	value float64
}

func (f Float) Value() float64 {
	return f.value
}

func NewFloat(value float64) Float {
	return Float{value: value}
}

// Refers to a literal boolean value
type Bool struct {
	isValueType
	isMethodIndependent
	value bool
}

func (b Bool) Value() bool {
	return b.value
}

func NewBool(value bool) Bool {
	return Bool{value: value}
}

// Refers to an array literal
type Array struct {
	isValueType
	elementType types.Any
	elements    []Any
}

func (a Array) ElementType() types.Any {
	return a.elementType
}

func (a Array) Elements() []Any {
	return a.elements
}

func (a Array) IsMethodDependent() bool {
	for _, element := range a.elements {
		if element.IsMethodDependent() {
			return true
		}
	}

	return false
}

func NewArray(elementType types.Any, element ...Any) Array {
	return Array{
		elementType: elementType,
		elements:    element,
	}
}

type KeyValue struct {
	key, value Any
}

func (k KeyValue) Key() Any {
	return k.key
}

func (k KeyValue) Value() Any {
	return k.value
}

func NewKeyValue(key, value Any) KeyValue {
	return KeyValue{
		key:   key,
		value: value,
	}
}

type Map struct {
	isValueType
	keyType, valueType types.Any
	elements           []KeyValue
}

func (m Map) KeyType() types.Any {
	return m.keyType
}

func (m Map) ValueType() types.Any {
	return m.valueType
}

func (m Map) Elements() []KeyValue {
	return m.elements
}

func (m Map) IsMethodDependent() bool {
	for _, element := range m.elements {
		if element.Key().IsMethodDependent() || element.Value().IsMethodDependent() {
			return true
		}
	}

	return false
}

func NewMap(keyType, valueType types.Any, elements []KeyValue) Map {
	return Map{
		keyType:   keyType,
		valueType: valueType,
		elements:  elements,
	}
}
