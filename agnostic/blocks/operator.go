package blocks

type MathOperator string

const (
	Add      MathOperator = "+"
	Subtract MathOperator = "-"
	Divide   MathOperator = "/"
	Multiply MathOperator = "*"
	Modulo   MathOperator = "%"
)

type ComparisonOperator string

const (
	Equal              ComparisonOperator = "=="
	NotEqual           ComparisonOperator = "!="
	GreatThan          ComparisonOperator = ">"
	GreatThanOrEqualTo ComparisonOperator = ">="
	LessThan           ComparisonOperator = "<"
	LassThanOrEqualTo  ComparisonOperator = "<="
)
