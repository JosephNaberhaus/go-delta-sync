package main

import (
	"errors"
	"flag"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
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

func GenerateImplementationTest(implementation blocks.Implementation) {
	implementation.Model("TestEmptyModel")
	implementation.Model(
		"TestModel",
		blocks.Field{Name: "IntField", TypeDescription: types.NewBaseTypeDescription(types.BaseTypeInt)},
	)

	testParameter := blocks.Field{Name: "testParameter", TypeDescription: types.NewBaseTypeDescription(types.BaseTypeInt)}
	implementation.Method("TestModel", "TestMethod", testParameter)
}
