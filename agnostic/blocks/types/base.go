package types

// All base types in Go. Most languages, besides Go, will probably support only a subset of these. Languages should be
// written to panic when that occurs.
type BaseType int

const (
	BaseTypeInt BaseType = iota
	BaseTypeInt8
	BaseTypeInt16
	BaseTypeInt32
	BaseTypeInt64
	BaseTypeUInt
	BaseTypeUInt8
	BaseTypeUInt16
	BaseTypeUInt32
	BaseTypeUInt64
	BaseTypeUIntPtr
	BaseTypeByte
	BaseTypeRune
	BaseTypeFloat32
	BaseTypeFloat64
	BaseTypeComplex64
	BaseTypeComplex128
)

var BaseTypeToGoValue = map[BaseType]string{
	BaseTypeInt:        "int",
	BaseTypeInt8:       "int8",
	BaseTypeInt16:      "int16",
	BaseTypeInt32:      "int32",
	BaseTypeInt64:      "int64",
	BaseTypeUInt:       "uint",
	BaseTypeUInt8:      "uint8",
	BaseTypeUInt16:     "uint16",
	BaseTypeUInt32:     "uint32",
	BaseTypeUInt64:     "uint64",
	BaseTypeUIntPtr:    "uintptr",
	BaseTypeByte:       "byte",
	BaseTypeRune:       "rune",
	BaseTypeFloat32:    "float32",
	BaseTypeFloat64:    "float64",
	BaseTypeComplex64:  "complex64",
	BaseTypeComplex128: "complex128",
}

type BaseTypeDescription struct {
	baseType BaseType
}

func NewBaseTypeDescription(baseType BaseType) BaseTypeDescription {
	return BaseTypeDescription{baseType: baseType}
}

func (b BaseTypeDescription) Classification() TypeClassification {
	return BaseClassification
}

func (b BaseTypeDescription) Value() string {
	return BaseTypeToGoValue[b.baseType]
}

func (b BaseTypeDescription) BaseType() BaseType {
	return b.baseType
}
