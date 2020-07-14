//go:generate go run ../../scripts/generate-test.go --impl go --implArg package:golang

package golang

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
	. "github.com/dave/jennifer/jen"
	"strings"
)

type Implementation struct {
	packageName string
	code        []Code
}

type BodyImplementation struct {
	receiverName string
	block        *Statement
}

func (g *Implementation) Add(c ...Code) {
	g.code = append(g.code, c...)
}

func (g *BodyImplementation) Add(c ...Code) {
	g.block.Add(lines(c...))
}

func (g *Implementation) Write(fileName string) {
	jenFile := NewFile(g.packageName)
	jenFile.Add(lines(g.code...))
	err := jenFile.Save(fileName + ".go")
	if err != nil {
		panic(err)
	}
}

func (g *Implementation) Model(modelName string, fields ...agnostic.Field) {
	modelStructFields := make([]Code, 0)
	for _, field := range fields {
		modelStructFields = append(modelStructFields, Id(field.Name).Add(resolveType(field.Type)))
	}

	g.Add(Type().Id(string(modelName)).Struct(modelStructFields...))
}

func (g *Implementation) Enum(name string, values ...string) {
	g.Add(Type().Id(name).Int())

	enumValues := make([]Code, 0)
	for i, v := range values {
		valueName := name + "_" + v
		if i == 0 {
			enumValues = append(enumValues, Id(valueName).Id(name).Op(":=").Iota())
		} else {
			enumValues = append(enumValues, Id(valueName))
		}
	}

	g.Add(Const().Defs(enumValues...))
}

func (g *Implementation) Method(modelName, methodName string, parameters ...agnostic.Field) agnostic.BodyImplementation {
	receiverName := strings.ToLower(modelName[:1])
	block := Null()

	parametersCode := make([]Code, 0)
	for _, param := range parameters {
		parametersCode = append(parametersCode, Id(param.Name).Add(resolveType(param.Type)))
	}

	g.Add(Func().Params(Id(receiverName).Op("*").Id(modelName)).Id(methodName).Params(parametersCode...).Block(block))

	return &BodyImplementation{
		receiverName: receiverName,
		block:        block,
	}
}

func (g *Implementation) ReturnMethod(modelName, methodName string, returnType types.Any, parameters ...agnostic.Field) agnostic.BodyImplementation {
	receiverName := strings.ToLower(modelName[:1])
	block := Null()

	parametersCode := make([]Code, 0)
	for _, param := range parameters {
		parametersCode = append(parametersCode, Id(param.Name).Add(resolveType(param.Type)))
	}

	g.Add(Func().Params(Id(receiverName).Op("*").Id(modelName)).Id(methodName).Params(parametersCode...).Add(resolveType(returnType)).Block(block))

	return &BodyImplementation{
		receiverName: receiverName,
		block:        block,
	}
}

func (g *BodyImplementation) Assign(assignee, assigned value.Any) {
	g.Add(resolveValue(assignee, g).Op("=").Add(resolveValue(assigned, g)))
}

func (g *BodyImplementation) Declare(name string, value value.Any) {
	g.Add(Id(name).Op(":=").Add(resolveValue(value, g)))
}

func (g *BodyImplementation) AppendValue(array, value value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(value, g)))
}

func (g *BodyImplementation) AppendArray(array, valueArray value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(valueArray, g).Op("...")))
}

func (g *BodyImplementation) RemoveValue(array, index value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(
		resolveValue(array, g).Index(Op(":").Add(resolveValue(index, g))),
		resolveValue(array, g).Index(resolveValue(index, g).Op("+").Lit(1).Op(":")).Op("..."),
	))
}

func (g *BodyImplementation) MapPut(mapValue, key, value value.Any) {
	g.Add(resolveValue(mapValue, g).Index(resolveValue(key, g)).Op("=").Add(resolveValue(value, g)))
}

func (g *BodyImplementation) MapDelete(mapValue, key value.Any) {
	g.Add(Delete(resolveValue(mapValue, g), resolveValue(key, g)))
}

func (g *BodyImplementation) ForEach(array value.Any, indexName, valueName string) agnostic.BodyImplementation {
	var forLoopParameter *Statement
	if indexName == "" {
		if valueName == "" {
			forLoopParameter = List()
		} else {
			forLoopParameter = List(Id("_"), Id(valueName)).Op(":=")
		}
	} else {
		if valueName == "" {
			forLoopParameter = List(Id(indexName)).Op(":=")
		} else {
			forLoopParameter = List(Id(indexName), Id(valueName)).Op(":=")
		}
	}

	block := Null()
	g.Add(For(forLoopParameter.Range().Add(resolveValue(array, g))).Block(block))
	return &BodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *BodyImplementation) If(value value.Any) agnostic.BodyImplementation {
	block := Null()
	g.Add(If(resolveValue(value, g)).Block(block))

	return &BodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *BodyImplementation) IfElse(value value.Any) (trueBody, falseBody agnostic.BodyImplementation) {
	trueBlock, falseBlock := Null(), Null()
	g.Add(If(resolveValue(value, g)).Block(trueBlock).Else().Block(falseBlock))

	trueBodyImplementation := &BodyImplementation{
		receiverName: g.receiverName,
		block:        trueBlock,
	}
	falseBodyImplementation := &BodyImplementation{
		receiverName: g.receiverName,
		block:        falseBlock,
	}

	return trueBodyImplementation, falseBodyImplementation
}

func (g *BodyImplementation) Return(value value.Any) {
	g.Add(Return(resolveValue(value, g)))
}

func NewImplementation(args map[string]string) agnostic.Implementation {
	packageName, ok := args["package"]
	if !ok {
		panic(errors.New("no package name supplied"))
	}

	return &Implementation{
		code:        make([]Code, 0),
		packageName: packageName,
	}
}

func resolveType(any types.Any) *Statement {
	switch t := any.(type) {
	case types.Base:
		return resolveBaseType(t)
	case types.Model:
		return Id(t.ModelName())
	case types.Array:
		return Index().Add(resolveType(t.Element()))
	case types.Map:
		return Map(resolveType(t.Key())).Add(resolveType(t.Value()))
	case types.Pointer:
		return Op("*").Add(resolveType(t.Value()))
	default:
		panic(errors.New(fmt.Sprintf("unkown type %T", t)))
	}
}

func resolveBaseType(base types.Base) *Statement {
	switch base {
	case types.BaseBool:
		return Bool()
	case types.BaseInt:
		return Int()
	case types.BaseInt32:
		return Int32()
	case types.BaseInt64:
		return Int64()
	case types.BaseFloat32:
		return Float32()
	case types.BaseFloat64:
		return Float64()
	case types.BaseString:
		return String()
	default:
		panic(errors.New("unknown base type " + string(base)))
	}
}

// Convert a value interface into its representation into Go code form
func resolveValue(any value.Any, optionalContext ...*BodyImplementation) *Statement {
	var context *BodyImplementation
	if len(optionalContext) == 0 && any.IsMethodDependent() {
		panic(errors.New("no context provided for method dependent value"))
	} else if len(optionalContext) == 1 {
		context = optionalContext[0]
	} else if len(optionalContext) > 1 {
		panic(errors.New("multiple body contexts provided when only 0 or 1 is allowed"))
	}

	switch v := any.(type) {
	case value.Null:
		return Nil()
	case value.String:
		return Lit(v.Value())
	case value.Int:
		return Lit(v.Value())
	case value.Float:
		return Lit(v.Value())
	case value.Array:
		elements := make([]Code, 0, len(v.Elements()))
		for _, element := range v.Elements() {
			elements = append(elements, resolveValue(element, context))
		}

		return Index().Add(resolveType(v.ElementType())).Values(elements...)
	case value.Map:
		elements := make([]Code, 0, len(v.Elements()))
		for _, element := range v.Elements() {
			elements = append(elements, resolveValue(element.Key()).Op(":").Add(resolveValue(element.Value())))
		}

		return Map(resolveType(v.KeyType())).Add(resolveType(v.ValueType())).Values(elements...)
	case value.OwnField:
		return Id(context.receiverName).Op(".").Add(resolveValue(v.Field(), context))
	case value.Id:
		return Id(v.Name())
	case value.ModelField:
		return Id(v.ModelName()).Op(".").Add(resolveValue(v.Field(), context))
	case value.ArrayElement:
		return resolveValue(v.Array(), context).Index(resolveValue(v.Index(), context))
	case value.MapElement:
		return resolveValue(v.Map(), context).Index(resolveValue(v.Key(), context))
	case value.Combined:
		return resolveValue(v.Left(), context).Op(v.Operator().Value()).Add(resolveValue(v.Right(), context))
	default:
		panic(errors.New(fmt.Sprintf("uknown type %T", v)))
	}
}

// Helper method to split a set of statements into lines of code
func lines(statements ...Code) Code {
	if len(statements) == 0 {
		return Null()
	}

	combined := Add(statements[0]).Line()
	for _, statement := range statements[1:] {
		combined = combined.Line().Add(statement)
	}

	return combined
}
