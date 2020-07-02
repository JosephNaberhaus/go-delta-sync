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
	Assign(assignee, assigned Value)
	ForEach(array Value, valueName string) BodyImplementation
	If(value1 Value, operator ComparisonOperator, value2 Value) BodyImplementation
}
