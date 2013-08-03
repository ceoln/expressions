package expressions

type copyVisitor int

func (copyVisitor) VisitConstant(value Float) (ExpressionLike, bool) {
	return NewConstant(value), true
}

func (copyVisitor) VisitVariableRef(s string) (ExpressionLike, bool) {
	return NewVariableRef(s), true
}

func (copyVisitor) VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool) {
	l, okay := lhs.DeepCopy()
	if !okay {
		return Zero, false
	}
	r, okay := rhs.DeepCopy()
	if !okay {
		return Zero, false
	}
	return NewBinaryOperation(op, l, r), true
}

/* DeepCopy returns an Expression with the same values, but sharing no actual data. */
func (e Expression) DeepCopy() (ExpressionLike, bool) {
	v := copyVisitor(0)
	answer, okay := e.Visit(v)
	if !okay {
		return Zero, false
	}
	return answer, true
}
