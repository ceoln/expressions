package expressions

import "testing"
import "fmt"

var LIMIT = 100

func TestRandom(t *testing.T) {
	var m = map[string]Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	var answer Float
	var ok bool
	var e Expression
	for i := 0; i < LIMIT; i++ {
		e = NewRandomExpression(m)
		answer, ok = e.Eval(m)
		if ok {
			break
		}
	}
	if ok == false {
		t.Errorf(fmt.Sprintf("Did not get an evaluatable random expression in %v tries.", LIMIT))
	} else {
		e2 := NewBinaryOperation('+', e, NewConstant(1))
		answer2, ok2 := e2.Eval(m)
		if ok2 == false {
			t.Errorf(fmt.Sprintf("Failed evaluation on %s.", e2))
		} else if answer2 != answer+1 {
			t.Errorf(fmt.Sprintf("%s evaluated to %v, but %s evaluated to %v, not %v+1", e, answer, e2, answer2, answer))
		}
	}
}
