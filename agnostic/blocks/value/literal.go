package value

import (
	"errors"
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

type Map struct {
	isValueType
	keyType, valueType types.Any
	keys, values       []Any
}

func (m Map) KeyType() types.Any {
	return m.keyType
}

func (m Map) ValueType() types.Any {
	return m.valueType
}

func (m Map) Keys() []Any {
	return m.keys
}

func (m Map) Values() []Any {
	return m.values
}

func (m Map) IsMethodDependent() bool {
	for _, key := range m.keys {
		if key.IsMethodDependent() {
			return true
		}
	}

	for _, value := range m.values {
		if value.IsMethodDependent() {
			return true
		}
	}

	return false
}

func NewMap(keyType, valueType types.Any, keys, values []Any) Map {
	if len(keys) != len(values) {
		panic(errors.New("cannot create map literal with mismatched number of keys and values"))
	}

	return Map{
		keyType:   keyType,
		valueType: valueType,
		keys:      keys,
		values:    values,
	}
}
