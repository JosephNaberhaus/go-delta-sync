package value

type Operator string

const (
	Add                Operator = "+"
	Subtract           Operator = "-"
	Divide             Operator = "/"
	Multiply           Operator = "*"
	Modulo             Operator = "%"
	Equal              Operator = "=="
	NotEqual           Operator = "!="
	GreatThan          Operator = ">"
	GreatThanOrEqualTo Operator = ">="
	LessThan           Operator = "<"
	LassThanOrEqualTo  Operator = "<="
)

func (m Operator) Value() string {
	return string(m)
}
