package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

var ForSuite = Suite{
	{
		Name:        "ForEachIndexAndElement",
		Description: "Support for foreach loop with both a value and index variable",
		ModelFields: []agnostic.Field{
			{Name: "SumElements", Type: types.BaseInt},
			{Name: "SumIndices", Type: types.BaseInt},
		},
		Parameters: []agnostic.Field{
			{Name: "arrayInput", Type: types.NewArray(types.BaseInt)},
		},
		Generator: func(body agnostic.BodyImplementation) {
			sumElementsValue := value.NewOwnField(value.NewId("SumElements"))
			sumIndicesValue := value.NewOwnField(value.NewId("SumIndices"))

			body.Assign(sumElementsValue, value.NewInt(0))
			body.Assign(sumIndicesValue, value.NewInt(0))

			forEachBody := body.ForEach(value.NewId("arrayInput"), "index", "value")
			forEachBody.Assign(sumElementsValue, value.NewCombined(sumElementsValue, value.Add, value.NewId("value")))
			forEachBody.Assign(sumIndicesValue, value.NewCombined(sumIndicesValue, value.Add, value.NewId("index")))
		},
		Facts: []Fact{
			{
				Name: "EmptyArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumElements",
						ExpectedValue: value.NewInt(0),
					},
					{
						FieldName:     "SumIndices",
						ExpectedValue: value.NewInt(0),
					},
				},
			},
			{
				Name: "PopulatedArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(10), value.NewInt(200)),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumElements",
						ExpectedValue: value.NewInt(211),
					},
					{
						FieldName:     "SumIndices",
						ExpectedValue: value.NewInt(3),
					},
				},
			},
		},
	},
	{
		Name:        "ForEachIndexOnly",
		Description: "Support for foreach loop with a index variable but no value variable",
		ModelFields: []agnostic.Field{
			{Name: "SumIndices", Type: types.BaseInt},
		},
		Parameters: []agnostic.Field{
			{Name: "arrayInput", Type: types.NewArray(types.BaseInt)},
		},
		Generator: func(body agnostic.BodyImplementation) {
			sumIndicesValue := value.NewOwnField(value.NewId("SumIndices"))

			body.Assign(sumIndicesValue, value.NewInt(0))

			forEachBody := body.ForEach(value.NewId("arrayInput"), "index", "")
			forEachBody.Assign(sumIndicesValue, value.NewCombined(sumIndicesValue, value.Add, value.NewId("index")))
		},
		Facts: []Fact{
			{
				Name: "EmptyArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumIndices",
						ExpectedValue: value.NewInt(0),
					},
				},
			},
			{
				Name: "PopulatedArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(10), value.NewInt(200)),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumIndices",
						ExpectedValue: value.NewInt(3),
					},
				},
			},
		},
	},
	{
		Name:        "ForEachElementOnly",
		Description: "Support for foreach loop with a value variable but no index variable",
		ModelFields: []agnostic.Field{
			{Name: "SumElements", Type: types.BaseInt},
		},
		Parameters: []agnostic.Field{
			{Name: "arrayInput", Type: types.NewArray(types.BaseInt)},
		},
		Generator: func(body agnostic.BodyImplementation) {
			sumElementsValue := value.NewOwnField(value.NewId("SumElements"))

			body.Assign(sumElementsValue, value.NewInt(0))

			forEachBody := body.ForEach(value.NewId("arrayInput"), "", "value")
			forEachBody.Assign(sumElementsValue, value.NewCombined(sumElementsValue, value.Add, value.NewId("value")))
		},
		Facts: []Fact{
			{
				Name: "EmptyArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumElements",
						ExpectedValue: value.NewInt(0),
					},
				},
			},
			{
				Name: "PopulatedArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(10), value.NewInt(200)),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "SumElements",
						ExpectedValue: value.NewInt(211),
					},
				},
			},
		},
	},
	{
		Name:        "ForEachNoLoopValues",
		Description: "Support for foreach loop with no value or index variables",
		ModelFields: []agnostic.Field{
			{Name: "NumElements", Type: types.BaseInt},
		},
		Parameters: []agnostic.Field{
			{Name: "arrayInput", Type: types.NewArray(types.BaseInt)},
		},
		Generator: func(body agnostic.BodyImplementation) {
			numElementsValue := value.NewOwnField(value.NewId("NumElements"))

			body.Assign(numElementsValue, value.NewInt(0))

			forEachBody := body.ForEach(value.NewId("arrayInput"), "", "")
			forEachBody.Assign(numElementsValue, value.NewCombined(numElementsValue, value.Add, value.NewInt(1)))
		},
		Facts: []Fact{
			{
				Name: "EmptyArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "NumElements",
						ExpectedValue: value.NewInt(0),
					},
				},
			},
			{
				Name: "PopulatedArray",
				Inputs: []value.Any{
					value.NewArray(types.BaseInt, value.NewInt(1), value.NewInt(10), value.NewInt(200)),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "NumElements",
						ExpectedValue: value.NewInt(3),
					},
				},
			},
		},
	},
}
