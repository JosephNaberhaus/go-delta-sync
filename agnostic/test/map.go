package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

var MapSuite = Suite{
	{
		Name:        "DeclareMap",
		Description: "Declares a map and then returns it",
		Returns:     types.NewMap(types.BaseInt, types.BaseString),
		Generator: func(body agnostic.BodyImplementation) {
			body.DeclareMap("declared", types.BaseInt, types.BaseString)
			body.Return(value.NewId("declared"))
		},
		Facts: []Fact{
			{
				Output: value.NewMap(types.BaseInt, types.BaseString),
			},
		},
	},
	{
		Name:        "MapPut",
		Description: "Puts a value into a map and returns it",
		Returns:     types.NewMap(types.BaseInt, types.BaseString),
		Parameters: []agnostic.Field{
			{Name: "mapInput", Type: types.NewMap(types.BaseInt, types.BaseString)},
			{Name: "key", Type: types.BaseInt},
			{Name: "value", Type: types.BaseString},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.MapPut(value.NewId("mapInput"), value.NewId("key"), value.NewId("value"))
			body.Return(value.NewId("mapInput"))
		},
		Facts: []Fact{
			{
				Name: "NewKey",
				Inputs: []value.Any{
					value.NewMap(types.BaseInt, types.BaseString),
					value.NewInt(1),
					value.NewString("test"),
				},
				SideEffects: nil,
				Output: value.NewMap(types.BaseInt, types.BaseString,
					value.NewKeyValue(value.NewInt(1), value.NewString("test")),
				),
			},
			{
				Name: "ExistingKey",
				Inputs: []value.Any{
					value.NewMap(types.BaseInt, types.BaseString,
						value.NewKeyValue(value.NewInt(1), value.NewString("initial")),
					),
					value.NewInt(1),
					value.NewString("new"),
				},
				SideEffects: nil,
				Output: value.NewMap(types.BaseInt, types.BaseString,
					value.NewKeyValue(value.NewInt(1), value.NewString("new")),
				),
			},
		},
	},
	{
		Name:        "MapDelete",
		Description: "Deletes a key from a map if it exists ",
		Returns:     types.NewMap(types.BaseInt, types.BaseString),
		Parameters: []agnostic.Field{
			{Name: "mapInput", Type: types.NewMap(types.BaseInt, types.BaseString)},
			{Name: "key", Type: types.BaseInt},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.MapDelete(value.NewId("mapInput"), value.NewId("key"))
			body.Return(value.NewId("mapInput"))
		},
		Facts: []Fact{
			{
				Name: "ExistentKey",
				Inputs: []value.Any{
					value.NewMap(types.BaseInt, types.BaseString,
						value.NewKeyValue(value.NewInt(1), value.NewString("test1")),
						value.NewKeyValue(value.NewInt(2), value.NewString("test2")),
					),
					value.NewInt(1),
				},
				SideEffects: nil,
				Output: value.NewMap(types.BaseInt, types.BaseString,
					value.NewKeyValue(value.NewInt(2), value.NewString("test2")),
				),
			},
			{
				Name: "NonExistentKey",
				Inputs: []value.Any{
					value.NewMap(types.BaseInt, types.BaseString,
						value.NewKeyValue(value.NewInt(2), value.NewString("test2")),
					),
					value.NewInt(1),
				},
				SideEffects: nil,
				Output: value.NewMap(types.BaseInt, types.BaseString,
					value.NewKeyValue(value.NewInt(2), value.NewString("test2")),
				),
			},
		},
	},
}
