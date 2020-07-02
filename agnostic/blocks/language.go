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
	Method(modelName, methodName string, parameters ...Field) (Implementation BodyImplementation)
}

type BodyImplementation interface {
	NewLine()
	Assign(assignee, assigned Value)
	ForEach(array Value, valueName string) (Implementation BodyImplementation)
	If(value1 Value, operator ComparisonOperator, value2 Value) (Implementation BodyImplementation)
}
