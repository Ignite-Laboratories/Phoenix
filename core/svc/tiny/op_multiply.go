package tiny

// Multiply multiplies each operand together in the order provided as a string Formula.
//
// see.BaselessCalculation
func Multiply(operands ...any) Formula[string] {
	return String.Multiply(operands...)
}

// Multiply multiplies each operand together in the order provided.
//
// see.BaselessCalculation
func (p Primitive[TOut]) Multiply(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Multiply multiplies each operand together in the order provided.
//
// see.BaselessCalculation
func (o Operation[TOut]) Multiply(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Multiply multiplies each operand together in the order provided.
//
// see.BaselessCalculation
func (f Formula[TOut]) Multiply(operands ...any) Formula[TOut] {
}
