//go:generate go run ../test/generate-test.go --impl go --implArg package:golang

package golang

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/values"
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

func (g *GoImplementation) Model(modelName blocks.ModelName, fields ...blocks.Field) {
	modelStructFields := make([]Code, 0)
	for _, field := range fields {
		modelStructFields = append(modelStructFields, Id(field.Name).Id(field.TypeDescription.Value()))
	}

	g.Add(Type().Id(string(modelName)).Struct(modelStructFields...))
}

func (g *GoImplementation) Method(modelName, methodName string, parameters ...blocks.Field) blocks.BodyImplementation {
	receiverName := strings.ToLower(modelName[:1])
	block := Null()

	parametersCode := make([]Code, 0)
	for _, param := range parameters {
		parametersCode = append(parametersCode, Id(param.Name).Id(param.TypeDescription.Value()))
	}

	g.Add(Func().Params(Id(receiverName).Op("*").Id(modelName)).Id(methodName).Params(parametersCode...).Block(block))

	return &GoBodyImplementation{
		receiverName: receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) Assign(assignee, assigned values.Any) {
	g.Add(resolveValue(assignee, g).Op("=").Add(resolveValue(assigned, g)))
}

func (g *GoBodyImplementation) Declare(name string, value values.Any) {
	g.Add(Id(name).Op(":=").Add(resolveValue(value, g)))
}

func (g *GoBodyImplementation) DeclareArray(name string, arrayType types.TypeDescription) {
	g.Add(Id(name).Op(":=").Make(Index().Id(arrayType.Value()), Lit(0)))
}

func (g *GoBodyImplementation) DeclareMap(name string, keyType, valueType types.TypeDescription) {
	g.Add(Id(name).Op(":=").Make(Map(Id(keyType.Value())).Id(valueType.Value())))
}

func (g *GoBodyImplementation) AppendValue(array, value values.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(value, g)))
}

func (g *GoBodyImplementation) AppendArray(array, valueArray values.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(resolveValue(array, g), resolveValue(valueArray, g).Op("...")))
}

func (g *GoBodyImplementation) RemoveValue(array, index values.Any) {
	g.Add(resolveValue(array, g).Op("=").Append(
		resolveValue(array, g).Index(Op(":").Add(resolveValue(index, g))),
		resolveValue(array, g).Index(resolveValue(index, g).Op("+").Lit(1).Op(":")),
	))
}

func (g *GoBodyImplementation) MapPut(mapValue, key, value values.Any) {
	g.Add(resolveValue(mapValue, g).Index(resolveValue(key, g)).Op("=").Add(resolveValue(value, g)))
}

func (g *GoBodyImplementation) MapDelete(mapValue, key values.Any) {
	g.Add(Delete(resolveValue(mapValue, g), resolveValue(key, g)))
}

func (g *GoBodyImplementation) ForEach(array values.Any, indexName, valueName string) blocks.BodyImplementation {
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

func (g *GoBodyImplementation) If(value values.Any) blocks.BodyImplementation {
	block := Null()
	g.Add(If(resolveValue(value, g)).Block(block))

	return &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) IfElse(value values.Any) (trueBody, falseBody blocks.BodyImplementation) {
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

func Implementation(args map[string]string) blocks.Implementation {
	packageName, ok := args["package"]
	if !ok {
		panic(errors.New("no package name supplied"))
	}

	return &GoImplementation{
		code:        make([]Code, 0),
		packageName: packageName,
	}
}

// Convert a value interface into its representation into Go code form
func resolveValue(value values.Any, context *GoBodyImplementation) *Statement {
	switch v := value.(type) {
	case values.Null:
		return Nil()
	case values.String:
		return Lit(v.Value())
	case values.Int:
		return Lit(v.Value())
	case values.Float:
		return Lit(v.Value())
	case values.OwnField:
		return Id(context.receiverName).Op(".").Add(resolveValue(v.Field(), context))
	case values.Id:
		return Id(v.Name())
	case values.ModelField:
		return Id(v.ModelName()).Op(".").Add(resolveValue(v.Field(), context))
	case values.Array:
		return resolveValue(v.Array(), context).Index(resolveValue(v.Index(), context))
	case values.Map:
		return resolveValue(v.Map(), context).Index(resolveValue(v.Key(), context))
	default:
		panic(errors.New("unknown value type " + fmt.Sprintf("%T", v)))
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
