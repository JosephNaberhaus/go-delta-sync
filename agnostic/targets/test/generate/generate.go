package generate

import (
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/types"
)

var TestSuites [][]Case = [][]Case{
	ArrayCases,
}

func GenerateAgnosticTests(implementation agnostic.Implementation) {
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
			Name: fmt.Sprintf("ArrayField%d", i),
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
		agnostic.Field{Name: "TestInt", Type: types.BaseInt},
		agnostic.Field{Name: "TestInt2", Type: types.BaseInt},
		agnostic.Field{Name: "TestArray", Type: types.NewArray(types.BaseInt)},
		agnostic.Field{Name: "TestMap", Type: types.NewMap(types.BaseInt, types.BaseInt)},
		agnostic.Field{Name: "IfOutput", Type: types.BaseString},
	)

	for _, suite := range TestSuites {
		for _, c := range suite {
			if c.Returns == nil {
				c.Generator(implementation.Method("TestModel", c.Name, c.Parameters...))
			} else {
				c.Generator(implementation.ReturnMethod("TestModel", c.Name, c.Returns, c.Parameters...))
			}
		}
	}
}

func GenerateImplementationTests(implementation Implementation) {
	for _, suite := range TestSuites {
		for _, c := range suite {
			implementation.Test(c)
		}
	}
}
