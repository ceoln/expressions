package expressions

import "math/rand"

/*
NewRandomExpression returns a random expression, using names from the given variable map (values are ignored).
Note that the expression may not evaluate successfully with a particular variable map, or
even with any variable map at all (if, for instance, it contains a hardcoded division by zero,
which it occasionally will).
*/
func NewRandomExpression(m map[string]Float) Expression {
	r := rand.Float64()
	switch {
	case r < 0.25:
		{
			return NewRandomConstant()
		}
	case r < 0.5:
		{
			return NewRandomVariableRef(m)
		}
	default:
		{
			return NewRandomBinOp(m)
		}
	}
}

/* NewRandomConstant returns a constant expression with some unpredictable value. */
func NewRandomConstant() Expression {
	return NewConstant(Float(float64(rand.Intn(100)) + rand.Float64()))
}

/* NewRandomVariableMap returns a variable reference to one of the variables in the map, uniformly distributed. */
func NewRandomVariableRef(m map[string]Float) Expression {
	var name string
	j := rand.Intn(len(m))
	for thisname := range m {
		if j == 0 {
			name = thisname
		}
		j--
	}
	return NewVariableRef(name)
}

/* NewRandomBinOp returns a binary operation expression in more or less the obvious way. */
func NewRandomBinOp(m map[string]Float) Expression {
	return NewBinaryOperation("+-*/"[rand.Intn(4)], NewRandomExpression(m), NewRandomExpression(m))
}
