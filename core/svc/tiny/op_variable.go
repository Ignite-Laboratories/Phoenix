package tiny

// NOTE: The variable system is used to build a list of named variables to print in LaTeX.
// This means an operation can name its own value, but a formula can name many operations.

// Variable simply stores the provided operand and yields a string Formula.
//
// see.BaselessCalculation
func Variable(name string, operand any) Formula[string] { return String.Variable(name, operand) }

// Variable simply stores the provided operand and yields a Formula.
//
// see.BaselessCalculation
func (o Operation[TOut]) Variable(name string, operand any) Formula[TOut] {
	return Formula[TOut]{
		op:   o,
		name: name,
		fn: func(Operation[TOut]) TOut {

		},
	}
}

// Variable wraps the entire Formula up into a single variable name.
//
// see.BaselessCalculation
func (f Formula[TOut]) Variable(name string) Formula[TOut] {
	f.name = name
	return f
}
