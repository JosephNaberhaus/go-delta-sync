//go:generate go run ../test/generate-test.go --impl go --implArg package:golang

package golang

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/block/value"
	. "github.com/dave/jennifer/jen"
	"strings"
)

type GoImplementation struct {
	packageName string
	code        []Code
}

type GoBodyImplementation struct {
	receiverName string
	block        *Statement
}

func (g *GoImplementation) Add(c ...Code) {
	g.code = append(g.code, c...)
}

func (g *GoBodyImplementation) Add(c ...Code) {
	g.block.Add(lines(c...))
}

func (g *GoImplementation) Write(fileName string) {
	jenFile := NewFile(g.packageName)
	jenFile.Add(lines(g.code...))
	err := jenFile.Save(fileName + ".go")
	if err != nil {
		panic(err)
	}
}

func (g *GoImplementation) Model(modelName agnostic.ModelName, fields ...agnostic.Field) {
	modelStructFields := make([]Code, 0)
	for _, field := range fields {
		modelStructFields = append(modelStructFields, Id(field.Name).Add(resolveType(field.Type)))
	}

	g.Add(Type().Id(string(modelName)).Struct(modelStructFields...))
}

func (g *GoImplementation) Method(modelName, methodName string, parameters ...agnostic.Field) agnostic.BodyImplementation {
	receiverName := strings.ToLower(modelName[:1])
	block := Null()

	parametersCode := make([]Code, 0)
	for _, param := range parameters {
		parametersCode = append(parametersCode, Id(param.Name).Add(resolveType(param.Type)))
	}

	g.Add(Func().Params(Id(receiverName).Op("*").Id(modelName)).Id(methodName).Params(parametersCode...).Block(block))

	return &GoBodyImplementation{
		receiverName: receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) Assign(assignee, assigned value.Any) {
	g.Add(resolveValue(assignee, g).Op("=").Add(resolveValue(assigned, g)))
}

func (g *GoBodyImplementation) Declare(name string, value value.Any) {
	g.Add(Id(name).Op(":=").Add(resolveValue(value, g)))
}

func (g *GoBodyImplementation) DeclareArray(name string, arrayType types.Any) {
	g.Add(Id(name).Op(":=").Make(Index().Add(resolveType(arrayType)), Lit(0)))
}

func (g *GoBodyImplementation) DeclareMap(name string, keyType, valueType types.Any) {
	g.Add(Id(name).Op(":=").Make(Map(Add(resolveType(keyType))).Add(resolveType(valueType))))
}

func (g *GoBodyImplementation) AppendValue(array, value value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(value, g)))
}

func (g *GoBodyImplementation) AppendArray(array, valueArray value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(valueArray, g).Op("...")))
}

func (g *GoBodyImplementation) RemoveValue(array, index value.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(
		resolveValue(array, g).Index(Op(":").Add(resolveValue(index, g))),
		resolveValue(array, g).Index(resolveValue(index, g).Op("+").Lit(1).Op(":")),
	))
}

func (g *GoBodyImplementation) MapPut(mapValue, key, value value.Any) {
	g.Add(resolveValue(mapValue, g).Index(resolveValue(key, g)).Op("=").Add(resolveValue(value, g)))
}

func (g *GoBodyImplementation) MapDelete(mapValue, key value.Any) {
	g.Add(Delete(resolveValue(mapValue, g), resolveValue(key, g)))
}

func (g *GoBodyImplementation) ForEach(array value.Any, indexName, valueName string) agnostic.BodyImplementation {
	if indexName == "" {
		indexName = "_"
	}

	if valueName == "" {
		indexName = "_"
	}

	block := Null()
	g.Add(For(Id(indexName), Id(valueName).Op(":=").Range().Add(resolveValue(array, g))).Block(block))
	return &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) If(value value.Any) agnostic.BodyImplementation {
	block := Null()
	g.Add(If(resolveValue(value, g)).Block(block))

	return &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) IfElse(value value.Any) (trueBody, falseBody agnostic.BodyImplementation) {
	trueBlock, falseBlock := Null(), Null()
	g.Add(If(resolveValue(value, g)).Block(trueBlock).Else().Block(falseBlock))

	trueBodyImplementation := &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        trueBlock,
	}
	falseBodyImplementation := &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        falseBlock,
	}

	return trueBodyImplementation, falseBodyImplementation
}

func Implementation(args map[string]string) agnostic.Implementation {
	packageName, ok := args["package"]
	if !ok {
		panic(errors.New("no package name supplied"))
	}

	return &GoImplementation{
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
func resolveValue(any value.Any, context *GoBodyImplementation) *Statement {
	switch v := any.(type) {
	case value.Null:
		return Nil()
	case value.String:
		return Lit(v.Value())
	case value.Int:
		return Lit(v.Value())
	case value.Float:
		return Lit(v.Value())
	case value.OwnField:
		return Id(context.receiverName).Op(".").Add(resolveValue(v.Field(), context))
	case value.Id:
		return Id(v.Name())
	case value.ModelField:
		return Id(v.ModelName()).Op(".").Add(resolveValue(v.Field(), context))
	case value.Array:
		return resolveValue(v.Array(), context).Index(resolveValue(v.Index(), context))
	case value.Map:
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
