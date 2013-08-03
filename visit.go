package expressions

/* Visitor is the interface for things that can Visit expressions (and more generally expressions.ExpressionLikes) */
type Visitor interface {
	VisitConstant(value Float) (ExpressionLike, bool)
	VisitVariableRef(s string) (ExpressionLike, bool)
	VisitBinaryOperation(op byte, lhs, rhs ExpressionLike) (ExpressionLike, bool)
}

/* Visit allows invoking a Visitor on each node of an Expression. */
func (e Expression) Visit(v Visitor) (ExpressionLike, bool) {
	switch e.op {
	case 'C':
		return v.VisitConstant(e.Value())
	case 'V':
		return v.VisitVariableRef(e.Name())
	default:
		lhs, rhs := e.LhsRhs()
		e1, okay := lhs.Visit(v)
		if !okay {
			return Zero, false
		}
		e2, okay := rhs.Visit(v)
		if !okay {
			return Zero, false
		}
		return v.VisitBinaryOperation(e.op, e1, e2)
	}
}
