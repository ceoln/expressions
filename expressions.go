package expressions

const VERSION = 1

/*
The type that is used for storing the scalar values that variables can take on, and that expressions evaluate to.
*/
type Float float64

/*
The basic type of an expression; obtained by the various NewXxx methods.
*/
type Expression struct {
	op    byte
	value Float  // Constant values, and accumulating Floats during visit walks
	name  string // Variable names, and accumulating strings during visit walks
	lhs   *ExpressionLike
	rhs   *ExpressionLike
}

/* Obtain a new constant expression, with the given numeric value. */
func NewConstant(c Float) Expression {
	return Expression{'C', c, "", nil, nil}
}

/* Obtain a new scalar variable reference, to a variable with the given name. */
func NewVariableRef(s string) Expression {
	return Expression{'V', 0, s, nil, nil}
}

/* Obtain a new binary expression, with the given operator ('+', '-', '*', or '/') and the given left and righthand sides. */
func NewBinaryOperation(op byte, e1 ExpressionLike, e2 ExpressionLike) Expression {
	return Expression{op, 0, "", &e1, &e2}
}

/* Test for equality; to anything! */
func (e Expression) Equal(f interface{}) bool {
	fe, okay := f.(ExpressionLike)
	if !okay {
		return false
	}
	return (e.String() == fe.String())
}

func (e Expression) Value() Float {
	return e.value
}

func (e Expression) Name() string {
	return e.name
}

func (e Expression) Op() byte {
	return e.op
}

func (e Expression) LhsRhs() (ExpressionLike, ExpressionLike) {
	return *e.lhs, *e.rhs
}

/* A useful obvious constant expression */
var Zero = NewConstant(0)
