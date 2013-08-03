package expressions

type evalVisitor map[string]Float // The object itself is the variable-value map

func (evalVisitor) VisitConstant(value Float) (ExpressionLike, bool) {
	return NewConstant(value), true
}

func (m evalVisitor) VisitVariableRef(s string) (ExpressionLike, bool) {
	v, okay := m[s]
	if !okay {
		return Zero, false
	}
	return NewConstant(v), true
}

func (m evalVisitor) VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool) {
	e1, ok := lhs.Eval(m)
	if ok == false {
		return Zero, false
	}
	e2, ok := rhs.Eval(m)
	if ok == false {
		return Zero, false
	}
	switch op {
	case '+':
		return NewConstant(e1 + e2), true
	case '-':
		return NewConstant(e1 - e2), true
	case '*':
		return NewConstant(e1 * e2), true
	case '/':
		if e2 == 0 {
			return Zero, false
		}
		return NewConstant(e1 / e2), true
	default:
		return Zero, false
	}
}

/*
Evaluate the ExpressionLike in the context of the given variable map, and return the result or a failure indication.
Failure results from referencing a variable not defined in the map, or from dividing by zero, or an unknown operation.
*/
func (e Expression) Eval(m map[string]Float) (Float, bool) {
	v := evalVisitor(m)
	answer, okay := e.Visit(v)
	if !okay {
		return 0, false
	}
	return answer.Value(), true
}

/* Eval without using a visitor, for testing and comparison and nostalgia. */
func (e Expression) EvalDirect(m map[string]Float) (Float, bool) {
	switch e.op {
	case 'C':
		return e.value, true
	case 'V':
		answer, ok := m[e.name]
		return answer, ok
	}
	lhs, okay := (*e.lhs).(Expression)
	if !okay {
		return 0, false
	}
	rhs, okay := (*e.rhs).(Expression)
	if !okay {
		return 0, false
	}
	e1, okay := lhs.EvalDirect(m)
	if !okay {
		return 0, false
	}
	e2, okay := rhs.EvalDirect(m)
	if !okay {
		return 0, false
	}
	switch e.op {
	case '+':
		return e1 + e2, true
	case '-':
		return e1 - e2, true
	case '*':
		return e1 * e2, true
	case '/':
		if e2 == 0 {
			return 0, false
		}
		return e1 / e2, true
	default:
		return 0, false
	}
}
