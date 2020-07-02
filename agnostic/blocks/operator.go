package blocks

type MathOperator string

const (
	Add      MathOperator = "+"
	Subtract MathOperator = "-"
	Divide   MathOperator = "/"
	Multiply MathOperator = "*"
	Modulo   MathOperator = "%"
)

func (m MathOperator) Value() string {
	return string(m)
}

type ComparisonOperator string

const (
	Equal              ComparisonOperator = "=="
	NotEqual           ComparisonOperator = "!="
	GreatThan          ComparisonOperator = ">"
	GreatThanOrEqualTo ComparisonOperator = ">="
	LessThan           ComparisonOperator = "<"
	LassThanOrEqualTo  ComparisonOperator = "<="
)

func (c ComparisonOperator) Value() string {
	return string(c)
}
