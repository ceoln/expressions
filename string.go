package expressions

import "fmt"

type stringVisitor int

const binaryOperationStringTemplate = "(%s%c%s)"

func (stringVisitor) VisitConstant(value Float) (ExpressionLike, bool) {
	return NewVariableRef(fmt.Sprintf("%v", value)), true
}

func (stringVisitor) VisitVariableRef(s string) (ExpressionLike, bool) {
	return NewVariableRef(s), true
}

func (stringVisitor) VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool) {
	return NewVariableRef(fmt.Sprintf(binaryOperationStringTemplate, lhs, op, rhs)), true
}

/* The string form of the expression (so that an Expression can be passed directly to fmt.print and family). */
func (e Expression) String() string {
	v := stringVisitor(0)
	answer, _ := e.Visit(v)
	return answer.Name()
}

/* Convert to string without using a visitor, for testing and comparison and so on. */
func (e Expression) StringDirect() string {
	switch e.op {
	case 'C':
		return fmt.Sprintf("%v", e.value)
	case 'V':
		return e.name
	default:
		return fmt.Sprintf(binaryOperationStringTemplate, (*e.lhs).String(), e.op, (*e.rhs).String())
	}
}
