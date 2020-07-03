package types

// All base types in Go. Most languages, besides Go, will probably support only a subset of these. Languages should be
// written to panic when that occurs.
type BaseType int

const (
	BaseTypeInt BaseType = iota
	BaseTypeInt32
	BaseTypeInt64
	BaseTypeFloat32
	BaseTypeFloat64
	BaseTypeBool
	BaseTypeString
)

var BaseTypeToGoValue = map[BaseType]string{
	BaseTypeInt:     "int",
	BaseTypeInt32:   "int32",
	BaseTypeInt64:   "int64",
	BaseTypeFloat32: "float32",
	BaseTypeFloat64: "float64",
	BaseTypeBool:    "bool",
	BaseTypeString:  "string",
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
