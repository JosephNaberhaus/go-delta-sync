//go:generate go run ../../scripts/generate-test.go --impl typescript

package typescript

import (
	"fmt"
	"io"
)

const IndentAmount = 2

type Code interface {
	Write(out io.Writer, indentLevel int) error
}

type Line string

func (n Line) Write(out io.Writer, indentLevel int) error {
	_, err := io.WriteString(out, fmt.Sprintf("%*s", indentLevel, n))
	return err
}

type Implementation struct {
	lines []Code
}

type BodyImplementation struct {
	lines []Code
}

func (i Implementation) Add(line Code) {
	i.lines = append(i.lines, line)
}

func NewBodyImplementation() BodyImplementation {
	return BodyImplementation{lines: make([]Code, 0)}
}

func (b BodyImplementation) Add(line Code) {
	b.lines = append(b.lines, line)
}

func (b BodyImplementation) Write(out io.Writer, indentLevel int) error {
	for _, line := range b.lines {
		err := line.Write(out, indentLevel+IndentAmount)
		if err != nil {
			return err
		}
	}

	return nil
}
