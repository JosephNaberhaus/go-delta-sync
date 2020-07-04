package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/value"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/targets"
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

	implementation, err := targets.CreateImplementation(implementationName, implementationArgs)
	if err != nil {
		panic(err)
	}

	GenerateImplementationTest(implementation)
	implementation.Write("agnostic-test")
}

func GenerateImplementationTest(implementation agnostic.Implementation) {
	implementation.Model("EmptyModel")

	// Create a model containing one of each base type
	baseTypeModelFields := make([]agnostic.Field, 0)
	for i := types.Base(0); i < types.NumberBaseTypes; i++ {
		newField := agnostic.Field{
			Name: fmt.Sprintf("Field%d", i),
			Type: i,
		}
		baseTypeModelFields = append(baseTypeModelFields, newField)
	}

	implementation.Model("BaseTypeModel", baseTypeModelFields...)

	// Create a model containing an array of each base type
	arrayTypeModelFields := make([]agnostic.Field, 0)
	for i := types.Base(0); i < types.NumberBaseTypes; i++ {
		newField := agnostic.Field{
			Name: fmt.Sprintf("ArrrayField%d", i),
			Type: types.NewArray(i),
		}
		arrayTypeModelFields = append(arrayTypeModelFields, newField)
	}

	implementation.Model("ArrayTypeModel", arrayTypeModelFields...)

	// Create a model containing a pointer of each base type
	pointerTypeModelFields := make([]agnostic.Field, 0)
	for i := types.Base(0); i < types.NumberBaseTypes; i++ {
		newField := agnostic.Field{
			Name: fmt.Sprintf("PointerField%d", i),
			Type: types.NewPointer(i),
		}
		pointerTypeModelFields = append(pointerTypeModelFields, newField)
	}

	implementation.Model("PointerTypeModel", pointerTypeModelFields...)

	// Create a model containing a map from each base type to every other base type
	mapTypeModelFields := make([]agnostic.Field, 0)
	for i := types.Base(0); i < types.NumberBaseTypes; i++ {
		for j := types.Base(0); j < types.NumberBaseTypes; j++ {
			newField := agnostic.Field{
				Name: fmt.Sprintf("MapField%dTo%d", i, j),
				Type: types.NewMap(i, j),
			}
			mapTypeModelFields = append(mapTypeModelFields, newField)
		}
	}

	implementation.Model("MapTypeModel", mapTypeModelFields...)

	// Create a model with as-is, pointer, array, and map variants of the empty model
	implementation.Model(
		"IdTypeModel",
		agnostic.Field{Name: "IdField", Type: types.NewModel("EmptyModel")},
		agnostic.Field{Name: "IdPointerField", Type: types.NewPointer(types.NewModel("EmptyModel"))},
		agnostic.Field{Name: "IdArrayField", Type: types.NewArray(types.NewModel("EmptyModel"))},
		agnostic.Field{Name: "IdToIdMapField", Type: types.NewMap(types.NewModel("EmptyModel"), types.NewModel("EmptyModel"))},
	)

	// Create a model to add test methods to
	implementation.Model(
		"TestModel",
		agnostic.Field{Name: "AssignInt", Type: types.BaseInt},
		agnostic.Field{Name: "IfOutput", Type: types.BaseString},
	)

	// Create test method that assigns the value to TestInteger
	parameter := agnostic.Field{Name: "value", Type: types.BaseInt}
	body := implementation.Method("TestModel", "TestAssign", parameter)
	body.Assign(value.NewOwnField(value.NewId("AssignInt")), value.NewId("value"))

	// Create a test method that declares a new variable that's the same as the
	// value and then assigns that to TestInteger
	parameter = agnostic.Field{Name: "value", Type: types.BaseInt}
	body = implementation.Method("TestModel", "TestDeclare", parameter)
	body.Declare("declared", value.NewId("value"))
	body.Assign(value.NewOwnField(value.NewId("AssignInt")), value.NewId("declared"))

	// Create a test method that sets IfOutput to "true" if the value is true
	parameter = agnostic.Field{Name: "value", Type: types.BaseBool}
	body = implementation.Method("TestModel ", "TestIf", parameter)
	body.If(value.NewId("value")).Assign(value.NewOwnField(value.NewId("IfOutput")), value.NewString("true"))

	// Create a test method that sets IfOutput to "true" if the value is true
	// and "false" otherwise
	parameter = agnostic.Field{Name: "value", Type: types.BaseBool}
	body = implementation.Method("TestModel", "TestIfElse", parameter)
	trueBody, falseBody := body.IfElse(value.NewId("value"))
	trueBody.Assign(value.NewOwnField(value.NewId("IfOutput")), value.NewString("true"))
	falseBody.Assign(value.NewOwnField(value.NewId("IfOutput")), value.NewString("false"))
}
