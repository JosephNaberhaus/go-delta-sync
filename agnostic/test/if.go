package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

var IfSuite = Suite{
	{
		Name:        "If",
		Description: "Assigns the value to the model if it is even or 0 otherwise",
		ModelFields: []agnostic.Field{
			{Name: "EvenNumber", Type: types.BaseInt},
		},
		Parameters: []agnostic.Field{
			{Name: "value", Type: types.BaseInt},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("EvenNumber")), value.NewInt(0))

			body.Declare("valueModulo2", value.NewCombined(value.NewId("value"), value.Modulo, value.NewInt(2)))
			ifBody := body.If(value.NewCombined(value.NewId("valueModulo2"), value.Equal, value.NewInt(0)))

			ifBody.Assign(value.NewOwnField(value.NewId("EvenNumber")), value.NewId("value"))
		},
		Facts: []Fact{
			{
				Name: "OddNumber",
				Inputs: []value.Any{
					value.NewInt(3),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "EvenNumber",
						ExpectedValue: value.NewInt(0),
					},
				},
			},
			{
				Name: "EvenNumber",
				Inputs: []value.Any{
					value.NewInt(4),
				},
				SideEffects: []SideEffect{
					{
						FieldName:     "EvenNumber",
						ExpectedValue: value.NewInt(4),
					},
				},
			},
		},
	},
	{
		Name:        "IfElse",
		Description: "Returns the string \"even\" if a value is even and the string \"odd\" otherwise",
		Parameters: []agnostic.Field{
			{Name: "value", Type: types.BaseInt},
		},
		Returns: types.BaseString,
		Generator: func(body agnostic.BodyImplementation) {
			body.Declare("valueModulo2", value.NewCombined(value.NewId("value"), value.Modulo, value.NewInt(2)))
			ifBody, elseBody := body.IfElse(value.NewCombined(value.NewId("valueModulo2"), value.Equal, value.NewInt(0)))

			ifBody.Return(value.NewString("even"))

			elseBody.Return(value.NewString("odd"))
		},
		Facts: []Fact{
			{
				Name: "OddNumber",
				Inputs: []value.Any{
					value.NewInt(3),
				},
				Output: value.NewString("odd"),
			},
			{
				Name: "EvenNumber",
				Inputs: []value.Any{
					value.NewInt(4),
				},
				Output: value.NewString("even"),
			},
		},
	},
}
