package tiny

// Factorial returns the result of "operand!" as a string Formula.
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func Factorial(operand any) Formula[string] {
	return String.Factorial(operand)
}

// Factorial returns the result of "operand!"
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func (p Primitive[TOut]) Factorial(operand any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Factorial returns the result of "operand!"
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func (o Operation[TOut]) Factorial(operand any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Factorial returns the result of "operand!"
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func (f Formula[TOut]) Factorial(operand any) Formula[TOut] {
}
