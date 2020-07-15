package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

var ArraySuite = Suite{
	{
		Name:        "AppendValue",
		Description: "Support for appending a value to an array",
		Parameters: []agnostic.Field{
			{Name: "inputArray", Type: types.NewArray(types.BaseInt)},
			{Name: "value", Type: types.BaseInt},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.AppendValue(value.NewId("inputArray"), value.NewId("value"))
			body.Return(value.NewId("inputArray"))
		},
		Facts: []Fact{
			{
				Name: "Empty",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
					value.NewInt(1),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1)),
			},
			{
				Name: "Populated",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewInt(4),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3), value.NewInt(4)),
			},
		},
	},
	{
		Name:        "AppendArray",
		Description: "Support for appending an array to another array",
		Parameters: []agnostic.Field{
			{Name: "inputArray", Type: types.NewArray(types.BaseInt)},
			{Name: "valueArray", Type: types.NewArray(types.BaseInt)},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.AppendArray(value.NewId("inputArray"), value.NewId("valueArray"))
			body.Return(value.NewId("inputArray"))
		},
		Facts: []Fact{
			{
				Name: "ModelEmpty",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
			},
			{
				Name: "ParameterEmpty",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewArray(types.BaseInt),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
			},
			{
				Name: "BothPopulated",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewArray(types.BaseInt, value.NewInt(4), value.NewInt(5), value.NewInt(6)),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3), value.NewInt(4), value.NewInt(5), value.NewInt(6)),
			},
		},
	},
	{
		Name:        "Remove",
		Description: "Support for removing an index from an array",
		Parameters: []agnostic.Field{
			{Name: "inputArray", Type: types.NewArray(types.BaseInt)},
			{Name: "index", Type: types.BaseInt},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.RemoveValue(value.NewId("inputArray"), value.NewId("index"))
			body.Return(value.NewId("inputArray"))
		},
		Facts: []Fact{
			{
				Name: "RemoveFirst",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewInt(0),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(2), value.NewInt(3)),
			},
			{
				Name: "RemoveMiddle",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewInt(1),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(3)),
			},
			{
				Name: "RemoveLast",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2), value.NewInt(3)),
					value.NewInt(2),
				},
				Output: value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(2)),
			},
		},
	},
}
