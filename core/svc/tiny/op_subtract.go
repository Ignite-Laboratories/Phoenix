package tiny

// Subtract subtracts each operand from the last in the order provided as a string Formula.
//
// see.BaselessCalculation
func Subtract(operands ...any) Formula[string] {
	return String.Subtract(operands...)
}

// Subtract subtracts each operand from the last in the order provided.
//
// see.BaselessCalculation
func (p Primitive[TOut]) Subtract(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Subtract subtracts each operand from the last in the order provided.
//
// see.BaselessCalculation
func (o Operation[TOut]) Subtract(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Subtract subtracts each operand from the last in the order provided.
//
// see.BaselessCalculation
func (f Formula[TOut]) Subtract(operands ...any) Formula[TOut] {
}
