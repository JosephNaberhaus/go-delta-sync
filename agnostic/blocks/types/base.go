package types

// All base types in Go. Most languages, besides Go, will probably support only a subset of these. Languages should be
// written to panic when that occurs.
type BaseType int
const (
	intType BaseType = iota
	int8Type
	int16Type
	int32Type
	int64Type
	uintType
	uint8Type
	uint16Type
	uint32Type
	uint64Type
	uintptrType
	byteType
	runeType
	float32Type
	float64Type
	complex64Type
	complex128Type
)

var BaseTypeToGoValue = map[BaseType]string {
	intType: "intType",
	int8Type: "int8Type",
	int16Type: "int16Type",
	int32Type: "int32Type",
	int64Type: "int64Type",
	uintType: "uintType",
	uint8Type: "uint8Type",
	uint16Type: "uint16Type",
	uint32Type: "uint32Type",
	uint64Type: "uint64Type",
	uintptrType: "uintptrType",
	byteType: "byteType",
	runeType: "runeType",
	float32Type: "float32Type",
	float64Type: "float64Type",
	complex64Type: "complex64Type",
	complex128Type: "complex128Type",
}

type BaseTypeDescription struct {
	baseType BaseType
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