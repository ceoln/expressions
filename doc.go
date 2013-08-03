/*
This expressions package defines a set of simple arithmetic expressions (on floating-point values),
that can be evaluated (in the context of a set of values for named variables), and printed out.

There is also a method to obtain a "random" expression, containing constants and variable names from
a given list.

Examples:
   var m = map[string] expressions.Float {
	"foo": 9,
	"bar": 10.5,
	"baz": 3.14159,
  }
  con := expressions.NewConstant(12.5)
  varble := expressions.NewVariableRef("foo")
  binop := expressions.NewBinaryOperation('+',con,varble)
  fmt.Println(binop)            // prints "12.5+foo"
  fmt.Println(binop.Eval(m))    // prints the answer, 21.5, and true (it worked)
  random := expressions.NewRandomExpression(m)
  answer,ok := random.Eval(m)
  fmt.Printf("%v: %v = %v\n",ok,random,answer)	// prints something random!

*/
package expressions
