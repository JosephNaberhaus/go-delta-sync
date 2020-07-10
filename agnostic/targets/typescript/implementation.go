//go:generate go run ../../scripts/generate-test.go --impl typescript

package typescript

import (
	"errors"
	"fmt"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/types"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/blocks/value"
	"io"
	"strconv"
	"strings"
)

const IndentAmount = 2

type Code interface {
	Write(out io.Writer, indentLevel int) error
}

type OrphanCode struct {
	belongsTo string
	code      []Code
}

func (o *OrphanCode) Add(code ...Code) {
	o.code = append(o.code, code...)
}

func NewOrphanCode(belongsTo string) *OrphanCode {
	return &OrphanCode{
		belongsTo: belongsTo,
		code:      make([]Code, 0),
	}
}

type Line string

func (n Line) Write(out io.Writer, indentLevel int) error {
	_, err := io.WriteString(out, fmt.Sprintf("%*s\n", indentLevel, n))
	return err
}

type Implementation struct {
	code        []Code
	modelBodies map[string]*BodyImplementation
	orphans     []*OrphanCode
}

func (i *Implementation) Add(code ...Code) {
	i.code = append(i.code, code...)
}

func (i *Implementation) AddOrphan(orphan *OrphanCode) {
	i.orphans = append(i.orphans, orphan)
}

type BodyImplementation struct {
	code []Code
}

func (b *BodyImplementation) Add(code ...Code) {
	b.code = append(b.code, code...)
}

func (b *BodyImplementation) Write(out io.Writer, indentLevel int) error {
	for _, line := range b.code {
		err := line.Write(out, indentLevel+IndentAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewBodyImplementation() *BodyImplementation {
	return &BodyImplementation{code: make([]Code, 0)}
}

func (i *Implementation) Model(name string, fields ...agnostic.Field) {
	body := NewBodyImplementation()
	for _, field := range fields {
		body.Add(Line(field.Name + ": " + resolveType(field.Type) + ";"))
	}

	i.Add(Line("export class " + name + "{"))
	i.Add(body)
	i.Add(Line("}"))
}
}

func resolveType(any types.Any) string {
	switch t := any.(type) {
	case types.Base:
		return resolveBaseType(t)
	case types.Model:
		return t.ModelName()
	case types.Array:
		return resolveType(t.Element()) + "[]"
	case types.Map:
		return "Map<" + resolveType(t.Key()) + ", " + resolveType(t.Value()) + ">"
	case types.Pointer:
		panic(errors.New("pointers are not supported yet"))
	default:
		panic(errors.New(fmt.Sprintf("unkown type %T", t)))
	}
}

func resolveBaseType(base types.Base) string {
	switch base {
	case types.BaseBool:
		return "boolean"
	case types.BaseInt:
		fallthrough
	case types.BaseInt32:
		fallthrough
	case types.BaseInt64:
		fallthrough
	case types.BaseFloat32:
		fallthrough
	case types.BaseFloat64:
		return "number"
	case types.BaseString:
		return "string"
	default:
		panic(errors.New("unknown base type " + string(base)))
	}
}

func resolveValue(any value.Any) string {
	switch v := any.(type) {
	case value.Null:
		return "null"
	case value.String:
		return "\"" + v.Value() + "\""
	case value.Int:
		return strconv.Itoa(v.Value())
	case value.Float:
		return strconv.FormatFloat(v.Value(), 'f', -1, 64)
	case value.Array:
		var sb strings.Builder

		sb.WriteString(resolveType(v.ElementType()) + "[] {")
		for i, element := range v.Elements() {
			sb.WriteString(resolveValue(element))

			if i-1 != len(v.Elements()) {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("}")

		return sb.String()
	case value.Map:
		var sb strings.Builder

		sb.WriteString(resolveType(types.NewMap(v.KeyType(), v.ValueType())) + "([")
		for i, element := range v.Elements() {
			sb.WriteString("[" + resolveValue(element.Key()) + ", " + resolveValue(element.Value()) + "]")

			if i-1 != len(v.Elements()) {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("])")

		return sb.String()
	case value.OwnField:
		return "this." + resolveValue(v.Field())
	case value.Id:
		return v.Name()
	case value.ModelField:
		return v.ModelName() + "." + resolveValue(v.Field())
	case value.ArrayElement:
		return resolveValue(v.Array()) + "[" + resolveValue(v.Index()) + "]"
	case value.MapElement:
		return resolveValue(v.Map()) + ".get(" + resolveValue(v.Key()) + ")"
	case value.Combined:
		return resolveValue(v.Left()) + " " + v.Operator().Value() + " " + resolveValue(v.Right())
	default:
		panic(errors.New(fmt.Sprintf("uknown type %T", v)))
	}
}
