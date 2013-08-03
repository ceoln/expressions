package expressions

import "testing"
import "fmt"

func confirm(e ExpressionLike, m map[string]Float, wantValue Float, wantOK bool, t *testing.T) {
	answer, ok := e.Eval(m)
	if ok != wantOK {
		t.Errorf("Eval(%v) = _,%v, want %v", e, ok, wantOK)
	} else if ok && (answer != wantValue) {
		t.Errorf("Eval(%v) = %v, want %v", e, answer, wantValue)
	}
}

func TestVersion(t *testing.T) {
	fmt.Println("Testing version",VERSION)
}

func TestConstant(t *testing.T) {
	var m map[string]Float = nil
	conExp := NewConstant(42)
	confirm(conExp, m, 42.0, true, t)
}

func TestVariable(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	varExp := NewVariableRef("foo")
	confirm(varExp, m, 7.0, true, t)
}

func TestArithmetic(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	confirm(NewBinaryOperation('+', NewConstant(9), NewConstant(12.5)), m, 9+12.5, true, t)
	confirm(NewBinaryOperation('-', NewConstant(9), NewConstant(12.5)), m, 9-12.5, true, t)
	confirm(NewBinaryOperation('*', NewConstant(9), NewConstant(12.5)), m, 9*12.5, true, t)
	confirm(NewBinaryOperation('/', NewConstant(9), NewConstant(12.5)), m, 9/12.5, true, t)
}

func TestNesting(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	confirm(NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo")), m, (5.0*200.0)+7.0, true, t)
}

func TestEvalDirect(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	want, _ := e.EvalDirect(m)
	confirm(e, m, want, true, t)
}

func TestString(t *testing.T) {
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	out := fmt.Sprintf("%s", e)
	want := "((5*200)+foo)"
	if out != want {
		t.Errorf("String(%v) = %v, want %v", e, out, want)
	}
}

func TestStringDirect(t *testing.T) {
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	want := fmt.Sprintf("%s", e)
	out := e.StringDirect()
	if out != want {
		t.Errorf("StringDirect(%v) = %v, want %v", e, out, want)
	}
}

func TestZero(t *testing.T) {
	confirm(Zero, nil, 0, true, t)
	if !Zero.Equal(NewConstant(0)) {
		t.Errorf("Zero.Equal(NewConstant(0)) not true.")
	}
}

func TestUnknownVar(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	varExp := NewVariableRef("qux")
	confirm(varExp, m, 0, false, t)
}

func TestDivisionByZero(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	confirm(NewBinaryOperation('/', NewConstant(8), NewConstant(0)), m, 0, false, t)
}
