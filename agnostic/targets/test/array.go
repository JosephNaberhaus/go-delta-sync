package main

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/value"
)

var ArrayCases = []Case{
	{
		Name:        "DeclareArray",
		Description: "Declares an array and then returns it",
		Generator: func(body agnostic.BodyImplementation) {
			body.DeclareArray("declared", types.BaseInt)
			body.Return(value.NewId("declared"))
		},
		Facts: []Fact{
			{
				Output: []int{},
			},
		},
	},
	{
		Name:        "AppendValue",
		Description: "Appends a value to an array and returns the result",
		Parameters: []agnostic.Field{
			{Name: "array", Type: types.NewArray(types.BaseInt)},
			{Name: "value", Type: types.BaseInt},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.AppendValue(value.NewId("array"), value.NewId("value"))
			body.Return(value.NewId("array"))
		},
		Facts: []Fact{
			{
				Inputs: []interface{}{
					[]int{},
					1,
				},
				Output: []int{1},
			},
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					4,
				},
				Output: []int{1, 2, 3, 4},
			},
		},
	},
	{
		Name:        "AppendArray",
		Description: "Appends a value to an array and returns the result",
		Parameters: []agnostic.Field{
			{Name: "array", Type: types.NewArray(types.BaseInt)},
			{Name: "valueArray", Type: types.NewArray(types.BaseInt)},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.AppendArray(value.NewId("array"), value.NewId("valueArray"))
			body.Return(value.NewId("array"))
		},
		Facts: []Fact{
			{
				Inputs: []interface{}{
					[]int{},
					[]int{1, 2, 3},
				},
				Output: []int{1, 2, 3},
			},
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					[]int{},
				},
				Output: []int{1, 2, 3},
			},
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
				},
				Output: []int{1, 2, 3, 4, 5, 6},
			},
		},
	},
	{
		Name:        "Remove",
		Description: "Removes a value from an array and returns the result",
		Parameters: []agnostic.Field{
			{Name: "array", Type: types.NewArray(types.BaseInt)},
			{Name: "index", Type: types.BaseInt},
		},
		Returns: types.NewArray(types.BaseInt),
		Generator: func(body agnostic.BodyImplementation) {
			body.RemoveValue(value.NewId("array"), value.NewId("index"))
			body.Return(value.NewId("array"))
		},
		Facts: []Fact{
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					0,
				},
				Output: []int{2, 3},
			},
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					1,
				},
				Output: []int{1, 3},
			},
			{
				Inputs: []interface{}{
					[]int{1, 2, 3},
					2,
				},
				Output: []int{1, 2},
			},
		},
	},
}
