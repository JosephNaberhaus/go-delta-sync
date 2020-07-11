package typescript

import (
	"bufio"
	"github.com/JosephNaberhaus/go-delta-sync/agnostic/test"
	"os"
	"strings"
)

const TestPreamble = `import {TestModel} from "./agnostic-test";
import * as assert from "assert";
describe('AgnosticTest', () => {
`

const TestPostscript = "});\n"

type TestImplementation struct {
	code           strings.Builder
	curIndentation int
}

func (t *TestImplementation) IncreaseIndentation() {
	t.curIndentation++
}

func (t *TestImplementation) DecreaseIndentation() {
	t.curIndentation--
}

func (t *TestImplementation) Add(line string) {
	t.code.WriteString(strings.Repeat("\t", t.curIndentation))
	t.code.WriteString(line)
	t.code.WriteString("\n")
}

func (t *TestImplementation) Write(fileName string) {
	file, err := os.Create(fileName + ".ts")
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(TestPreamble)
	if err != nil {
		panic(err)
	}

	_, err = writer.WriteString(t.code.String())
	if err != nil {
		panic(err)
	}

	_, err = writer.WriteString(TestPostscript)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func (t *TestImplementation) Test(testCase test.Case) {
	for _, fact := range testCase.Facts {
		t.Add("it('" + testCase.Name + "_" + fact.Name + " should work', () => {")
		t.IncreaseIndentation()

		t.Add("const model = new TestModel();")

		var inputs strings.Builder
		for i, input := range fact.Inputs {
			inputs.WriteString(resolveValue(input))

			if i+1 != len(fact.Inputs) {
				inputs.WriteString(", ")
			}
		}

		if fact.Output != nil {
			t.Add("const result = model." + testCase.Name + "(" + inputs.String() + ");")
			t.Add("assert.deepStrictEqual(" + resolveValue(fact.Output) + ", result);")
		} else {
			t.Add("model." + testCase.Name + "(" + inputs.String() + ");")
		}

		for _, sideEffect := range fact.SideEffects {
			t.Add("assert.deepStrictEqual(" + resolveValue(sideEffect.ExpectedValue) + ", model." + sideEffect.FieldName + ");")
		}

		t.DecreaseIndentation()
		t.Add("});")
	}
}

func NewTestImplementation(args map[string]string) test.Implementation {
	return &TestImplementation{
		curIndentation: 1,
	}
}
