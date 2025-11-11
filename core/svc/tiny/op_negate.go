package tiny

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func Negate(operand any) Formula[string] {
	return String.Negate(operand)
}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func (p Primitive[TOut]) Negate(operand any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func (o Operation[TOut]) Negate(operand any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func (f Formula[TOut]) Negate(operand any) Formula[TOut] {
}
