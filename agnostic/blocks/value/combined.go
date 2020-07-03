package value

type Combined struct {
	valueType
	left, right Any
	operator    Operator
}

func (c Combined) Left() Any {
	return c.left
}

func (c Combined) Right() Any {
	return c.right
}

func (c Combined) Operator() Operator {
	return c.operator
}

func NewCombined(left Any, operator Operator, right Any) Combined {
	return Combined{
		left:     left,
		right:    right,
		operator: operator,
	}
}
