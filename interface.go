package expressions

import "fmt"

/* ExpressionLike is the interface for things that implement what an Expression does. */
type ExpressionLike interface {
	fmt.Stringer
	Value() Float
	Name() string
	Op() byte
	LhsRhs() (ExpressionLike, ExpressionLike)
	Equal(f interface{}) bool
	DeepCopy() (ExpressionLike, bool)
	Visit(v Visitor) (ExpressionLike, bool)
	Eval(m map[string]Float) (Float, bool)
}
