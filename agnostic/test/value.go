package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

// Test value types. Most of these tests are written with the intention of
// catching syntax errors in generated code. Since the value generation code
// is the same code that is used for test values, it is is difficult to
// actually test that Agnostic is doing what is expected. To get around that,
// arithmetic can be used to represent the same number in multiple ways.
var ValueSuite = Suite{
	{
		Name:        "Int",
		Description: "Assigns models fields to various int type values",
		ModelFields: []agnostic.Field{
			{Name: "Int1", Type: types.BaseInt},
			{Name: "Int2", Type: types.BaseInt},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("Int1")), value.NewInt(-100))
			body.Assign(value.NewOwnField(value.NewId("Int2")), value.NewInt(100))
		},
		Facts: []Fact{
			{
				Name: "AssignsValues",
				SideEffects: []SideEffect{
					{FieldName: "Int1", ExpectedValue: value.NewInt(-100)},
					{FieldName: "Int2", ExpectedValue: value.NewInt(100)},
				},
			},
		},
	},
	{
		Name:        "IntPastMaxSafeFloatInteger",
		Description: "Assigns models fields to various int type values that are past what can safely be represented in a IEE-754 floating point number",
		ModelFields: []agnostic.Field{
			{Name: "Int1", Type: types.BaseInt},
			{Name: "Int2", Type: types.BaseInt},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("Int1")), value.NewInt(-16777216))
			body.Assign(value.NewOwnField(value.NewId("Int2")), value.NewInt(16777216))
		},
		Facts: []Fact{
			{
				Name: "AssignsValues",
				SideEffects: []SideEffect{
					{FieldName: "Int1", ExpectedValue: value.NewCombined(value.NewInt(-16777217), value.Add, value.NewInt(1))},
					{FieldName: "Int2", ExpectedValue: value.NewCombined(value.NewInt(16777217), value.Subtract, value.NewInt(1))},
				},
			},
		},
	},
	{
		Name:        "IntPastMaxSafeDoubleInteger",
		Description: "Assigns models fields to various int type values that are past what can safely be represented in a double precision IEE-754 floating point number",
		ModelFields: []agnostic.Field{
			{Name: "Int1", Type: types.BaseInt},
			{Name: "Int2", Type: types.BaseInt},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("Int1")), value.NewInt(-9007199254740992))
			body.Assign(value.NewOwnField(value.NewId("Int2")), value.NewInt(9007199254740992))
		},
		Facts: []Fact{
			{
				Name: "AssignsValues",
				SideEffects: []SideEffect{
					{FieldName: "Int1", ExpectedValue: value.NewCombined(value.NewInt(-9007199254740993), value.Add, value.NewInt(1))},
					{FieldName: "Int2", ExpectedValue: value.NewCombined(value.NewInt(9007199254740993), value.Subtract, value.NewInt(1))},
				},
			},
		},
	},
}
