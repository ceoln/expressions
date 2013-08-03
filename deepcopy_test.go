package expressions

import "testing"

func TestBasicCopy(t *testing.T) {
	e := NewBinaryOperation('+', NewBinaryOperation('*', NewConstant(5), NewConstant(200)), NewVariableRef("foo"))
	f, okay := e.DeepCopy()
	if !okay {
		t.Errorf("%v copy result was not okay", e)
	}
	if e.String() != f.String() {
		t.Errorf("%v copy result was %v, want %v", e, f, e)
	}
}
