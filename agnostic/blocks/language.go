package blocks

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
)

type ModelName string

type Field struct {
	Name            string
	TypeDescription types.TypeDescription
}

type Implementation interface {
	Write(fileName string)
	Model(ModelName, ...Field)
	Method(modelName, methodName string, parameters ...Field) BodyImplementation
}

type BodyImplementation interface {
	// Assigns the value at assigned to assignee
	// Go Code: `<assignee> = <assigned>`
	Assign(assignee, assigned Value)
	// Declares a new variable with the given value
	// Go Code: `<declared> := <value>`
	Declare(declared VariableStruct, value Value)

	// Iterates through every value of the given array. Index name and value
	// are to be variables containing the equivalent of a zero based index and
	// the value at that index. An empty string for a name will  to indicate to
	// the implementation that the value is not used
	// Go Code: `for <indexName>, <valueName> := range <array> { <body> }
	ForEach(array Value, indexName, valueName string) BodyImplementation

	// Performs a comparison operation on the two values and executed the body
	// if the result is true
	// Go Code: `if <value1> <operator> <value2> { <body> }
	If(value1 Value, operator ComparisonOperator, value2 Value) BodyImplementation
	// Performs a comparison operation on the two values and executes the true
	// body if the result is true and the false body otherwise
	// Go Code: `if <value1> <operator> <value2> { <true block> } else { <false block> }
	IfElse(value1 Value, operator ComparisonOperator, value2 Value) (TrueBody, FalseBody BodyImplementation)
	// Executes the body if the value is true
	// Go Code: `if <value> { <body> }
	IfBool(value Value) BodyImplementation
	// Execute the true body if the value is true and the false body otherwise
	// Go Code: `if <value> { <true body> } else { <false body> }
	IfElseBool(value Value) (TrueBody, FalseBody BodyImplementation)
}
