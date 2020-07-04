package blocks

import (
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
)

type ModelName string

type Field struct {
	Name string
	Type types.Any
}

// A file in an arbitrary programming language
type Implementation interface {
	// Writes the current contents of the file to the given path. The path
	// should exclude the extension which which will be added by the
	// implementation
	Write(fileName string)

	// Creates a new model
	// Go Code: type <name> struct { <fields> }
	Model(name ModelName, fields ...Field)

	// Create a new method. A method is simply a function that runs under the
	// context of a model and has direct access to it's contents
	// Go Code: func (<first character of modelName> *<modelName>) <methodName>(<parameters>) { <body> }
	Method(modelName, methodName string, parameters ...Field) BodyImplementation
}

// An ordered set of logic that runs inside of a method.
type BodyImplementation interface {
	// Assigns the value at assigned to assignee
	// Go Code: `<assignee> = <assigned>`
	Assign(assignee, assigned value.Any)
	// Declares a new variable with the given value
	// Go Code: `<name> := <value>`
	Declare(name string, value value.Any)
	// Declares a new array variable containing value of the given type
	// Go Code: `<name> := make(<arrayType>[]>, 0)
	DeclareArray(name string, arrayType types.Any)
	// Declares a new map variable
	// Go Code: `<name> := make(map[<keyType>]<valueType>)
	DeclareMap(name string, keyType, valueType types.Any)

	// Appends a value to the end of an array and ensures that the array value
	// points to the result. This comes with no guarantees that a different
	// reference to the array will not also be modified
	// Go Code: '<array> = append(<array>, <value>)`
	AppendValue(array, value value.Any)
	// Appends an array to the end of another array and ensures that the array
	// value points to the result. This comes with no guarantee that a
	// different reference to the array that was appended to will not also be
	// modified. However, the value array will not be altered
	// Go Code: `<array> = append(<array>, <valueArray>...)`
	AppendArray(array, valueArray value.Any)
	// Remove the value at index from the array. The order of the array must
	// not be altered by this operation and it must leave no gap from where the
	// element was removed.
	// Go Code: `<array> = append<array[:<index>], <array>[<index>+1:]...)`
	RemoveValue(array, index value.Any)

	// Sets key to value in the map, overriding an existing value or creating a
	// new entry a necessary
	// Go Code: `<mapValue>[<key>] = <value>`
	MapPut(mapValue, key, value value.Any)
	// Deleted the given value from the map. However this is performed the key
	// must be considered to no longer exist on the map
	// Go Code: `delete(<mapValue>, <key>)`
	MapDelete(mapValue, key value.Any)

	// Iterates through every value of the given array. Index name and value
	// are to be variables containing the equivalent of a zero based index and
	// the value at that index. An empty string for a name will  to indicate to
	// the implementation that the value is not used
	// Go Code: `for <indexName>, <valueName> := range <array> { <body> }
	ForEach(array value.Any, indexName, valueName string) BodyImplementation

	// Executes the body if the value is true
	// Go Code: `if <value> { <body> }
	If(value value.Any) BodyImplementation
	// Execute the true body if the value is true and the false body otherwise
	// Go Code: `if <value> { <true body> } else { <false body> }
	IfElse(value value.Any) (TrueBody, FalseBody BodyImplementation)
}
