package main

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/types"
)

// A function that takes the given body implementation and creates a test case
// inside of it
type GenerateFunc func(body agnostic.BodyImplementation)

// A method should be created
type Case struct {
	Name        string           // Name of the test case (must be unique)
	Description string           // Describes what the test case is for
	Parameters  []agnostic.Field // Parameters that the generated test method will take in
	Returns     types.Any        // The return type of the method or nil if it returns nothing
	Generator   GenerateFunc     // Function that generates the test method
	Facts       []Fact           // Facts about the Test
}

// A change that happens to the model as a result of a method call
type SideEffect struct {
	FieldName     string      // Name of the field
	ExpectedValue interface{} // The expected value of the field
}

func NewSideEffect(fieldName string, expectedValue interface{}) SideEffect {
	return SideEffect{
		FieldName:     fieldName,
		ExpectedValue: expectedValue,
	}
}

// An assertion that calling the method with the given inputs will result in
// the given side effects and output
type Fact struct {
	Inputs      []interface{} // Values to pass in a parameters or nil
	SideEffects []SideEffect  // Side effects of the method call or nil
	Output      interface{}   // The output of the method or nil
}
