package tiny

// Absolute returns the result of "|operand|" as a string Formula.
//
// see.BaselessCalculation
func Absolute(operand any) Formula[string] {
	return String.Absolute(operand)
}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func (p Primitive[TOut]) Absolute(operand any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func (o Operation[TOut]) Absolute(operand any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func (f Formula[TOut]) Absolute(operand any) Formula[TOut] {
}
