package main

import (
	"errors"
	"flag"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	. "github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks"
	. "github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"strings"
)

type flagImplementationArgs map[string]string

func (i flagImplementationArgs) Set(value string) error {
	if len(value) == 0 {
		return nil
	}

	split := strings.Split(value, ":")
	if len(split) != 2 {
		return errors.New("implementation arg must be in form <key>:<value>")
	}

	i[split[0]] = split[1]
	return nil
}

func (i flagImplementationArgs) String() string {
	var sb strings.Builder
	for key, value := range i {
		sb.WriteString(key)
		sb.WriteString(":")
		sb.WriteString(value)
		sb.WriteString(" ")
	}

	return sb.String()
}

func main() {
	var implementationName string
	var implementationArgs = make(flagImplementationArgs)

	flag.StringVar(&implementationName, "impl", "", "language implementation name/path to use")
	flag.Var(&implementationArgs, "implArg", "arguments to pass into implementation constructor")

	flag.Parse()

	if len(implementationName) == 0 {
		panic(errors.New("implementation name/path is required"))
	}

	implementation, err := agnostic.CreateImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	GenerateImplementationTest(implementation)
	implementation.Write("agnostic-test")
}

func GenerateImplementationTest(implementation Implementation) {
	implementation.Model("TestEmptyModel")
	implementation.Model(
		"TestModel",
		Field{Name: "IntField", TypeDescription: NewBaseTypeDescription(BaseTypeInt)},
		Field{Name: "Int8Field", TypeDescription: NewBaseTypeDescription(BaseTypeInt8)},
		Field{Name: "Int16Field", TypeDescription: NewBaseTypeDescription(BaseTypeInt16)},
		Field{Name: "Int32Field", TypeDescription: NewBaseTypeDescription(BaseTypeInt32)},
		Field{Name: "Int64Field", TypeDescription: NewBaseTypeDescription(BaseTypeInt64)},
		Field{Name: "UIntField", TypeDescription: NewBaseTypeDescription(BaseTypeUInt)},
		Field{Name: "UInt8Field", TypeDescription: NewBaseTypeDescription(BaseTypeUInt8)},
		Field{Name: "UInt16Field", TypeDescription: NewBaseTypeDescription(BaseTypeUInt16)},
		Field{Name: "UInt32Field", TypeDescription: NewBaseTypeDescription(BaseTypeUInt32)},
		Field{Name: "UInt64Field", TypeDescription: NewBaseTypeDescription(BaseTypeUInt64)},
		Field{Name: "UIntPtrField", TypeDescription: NewBaseTypeDescription(BaseTypeUIntPtr)},
		Field{Name: "ByteField", TypeDescription: NewBaseTypeDescription(BaseTypeByte)},
		Field{Name: "RuneField", TypeDescription: NewBaseTypeDescription(BaseTypeRune)},
		Field{Name: "Float32Field", TypeDescription: NewBaseTypeDescription(BaseTypeFloat32)},
		Field{Name: "Float64Field", TypeDescription: NewBaseTypeDescription(BaseTypeFloat64)},
		Field{Name: "Complex64Field", TypeDescription: NewBaseTypeDescription(BaseTypeComplex64)},
		Field{Name: "Complex128Field", TypeDescription: NewBaseTypeDescription(BaseTypeComplex128)},
	)

	testParameter := Field{Name: "testParameter", TypeDescription: NewBaseTypeDescription(BaseTypeInt)}
	implementation.Method("TestModel", "TestMethod", testParameter)
}
