package test

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

var ValueSuite = Suite{
	{
		Name:        "Int",
		Description: "Support for integer types",
		ModelFields: []agnostic.Field{
			{Name: "String1", Type: types.BaseString},
			{Name: "String2", Type: types.BaseString},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("String1")), value.NewIntToString(value.NewInt(-100)))
			body.Assign(value.NewOwnField(value.NewId("String2")), value.NewIntToString(value.NewInt(100)))
		},
		Facts: []Fact{
			{
				Name: "StringValueMatches",
				SideEffects: []SideEffect{
					{FieldName: "String1", ExpectedValue: value.NewString("-100")},
					{FieldName: "String2", ExpectedValue: value.NewString("100")},
				},
			},
		},
	},
	{
		Name:        "IntPastMaxSafeFloatInteger",
		Description: "Support for numbers outside of the range [-16777216, 16777216]",
		ModelFields: []agnostic.Field{
			{Name: "String1", Type: types.BaseString},
			{Name: "String2", Type: types.BaseString},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("String1")), value.NewIntToString(value.NewInt(-16777217)))
			body.Assign(value.NewOwnField(value.NewId("String2")), value.NewIntToString(value.NewInt(16777217)))
		},
		Facts: []Fact{
			{
				Name: "StringValueMatches",
				SideEffects: []SideEffect{
					{FieldName: "String1", ExpectedValue: value.NewString("-16777217")},
					{FieldName: "String2", ExpectedValue: value.NewString("16777217")},
				},
			},
		},
	},
	{
		Name:        "IntPastMaxSafeDoubleInteger",
		Description: "Support for numbers outside of the range [-9007199254740992, 9007199254740992]",
		ModelFields: []agnostic.Field{
			{Name: "String1", Type: types.BaseString},
			{Name: "String2", Type: types.BaseString},
		},
		Generator: func(body agnostic.BodyImplementation) {
			body.Assign(value.NewOwnField(value.NewId("String1")), value.NewIntToString(value.NewInt(-9007199254740993)))
			body.Assign(value.NewOwnField(value.NewId("String2")), value.NewIntToString(value.NewInt(9007199254740993)))
		},
		Facts: []Fact{
			{
				Name: "StringValueMatches",
				SideEffects: []SideEffect{
					{FieldName: "String1", ExpectedValue: value.NewString("-9007199254740993")},
					{FieldName: "String2", ExpectedValue: value.NewString("9007199254740993")},
				},
			},
		},
	},
}
