//go:generate go run ../test/generate-test.go --output implementation-test.go --impl go --implArg package:golang

package golang

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks"
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
	g.code = append(g.code, lines(c...))
}

func (g *GoBodyImplementation) Add(c ...Code) {
	g.block.Add(lines(c...))
}

func (g *GoImplementation) Write(fileName string) {
	jenFile := NewFile(g.packageName)
	jenFile.Add(g.code...)
	err := jenFile.Save(fileName)
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
		receiverName: modelName[:1],
		block:        block,
	}
}

func (g *GoBodyImplementation) Assign(assignee, assigned blocks.Value) {
	g.Add(resolveValue(assignee, g).Op("=").Add(resolveValue(assigned, g)))
}

func (g *GoBodyImplementation) ForEach(array blocks.Value, indexName, valueName string) blocks.BodyImplementation {
	block := Null()
	g.Add(For(Id(indexName), Id(valueName).Op(":=").Range().Add(resolveValue(array, g))).Block(block))
	return &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
}

func (g *GoBodyImplementation) If(value1 blocks.Value, operator blocks.ComparisonOperator, value2 blocks.Value) blocks.BodyImplementation {
	block := Null()
	g.Add(If(resolveValue(value1, g).Op(operator.Value()).Add(resolveValue(value2, g))).Block(block))
	return &GoBodyImplementation{
		receiverName: g.receiverName,
		block:        block,
	}
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
func resolveValue(value blocks.Value, context *GoBodyImplementation) *Statement {
	switch v := value.(type) {
	case blocks.NullValue:
		return Nil()
	case blocks.StringValue:
		return Lit(v.Value())
	case blocks.IntValue:
		return Lit(v.Value())
	case blocks.FloatValue:
		return Lit(v.Value())
	case blocks.OwnProperty:
		return Id(context.receiverName).Dot(v.Name())
	case blocks.Variable:
		return Id(v.Name())
	case blocks.ModelProperty:
		return Id(v.ModelName()).Dot(v.Name())
	case blocks.ArrayValue:
		return resolveValue(v.Array(), context).Index(resolveValue(v.Index(), context))
	case blocks.MapValue:
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

	combined := Null()
	for _, statement := range statements {
		combined = combined.Line().Add(statement)
	}

	return combined
}
