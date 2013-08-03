package expressions

import "testing"
import "fmt"

/* visitor that converts to string without parens */

type visitTester1 int

func (visitTester1) VisitConstant(value Float) (ExpressionLike, bool) {
	return NewVariableRef(fmt.Sprintf("%v", value)), true
}

func (visitTester1) VisitVariableRef(s string) (ExpressionLike, bool) {
	return NewVariableRef(s), true
}

func (visitTester1) VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool) {
	return NewVariableRef(fmt.Sprintf("%s%c%s", lhs, op, rhs)), true
}

func TestVisit1(t *testing.T) {
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	v := visitTester1(0)
	answerv, okay := e.Visit(v)
	if !okay {
		t.Errorf("%v visit1 result was not okay", e)
	}
	answer := answerv.(fmt.Stringer)
	out := answer.String()
	want := "5*200+foo"
	if out != want {
		t.Errorf("%v visit1 result was %v, want %v", e, out, want)
	}
}

/* visitor that just counts leaf nodes, as a pointer receiver */

type visitTester2 int

func (v *visitTester2) VisitConstant(value Float) (ExpressionLike, bool) {
	*v++
	return Zero, true
}

func (v *visitTester2) VisitVariableRef(s string) (ExpressionLike, bool) {
	*v++
	return Zero, true
}

func (*visitTester2) VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool) {
	return Zero, true
}

func TestVisit2(t *testing.T) {
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	v := visitTester2(0)
	_, okay := e.Visit(&v)
	if !okay {
		t.Errorf("%v visit2 result was not okay", e)
	}
	if v != 3 {
		t.Errorf("%v visit2 count was %v, want %v", e, v, 3)
	}
}
