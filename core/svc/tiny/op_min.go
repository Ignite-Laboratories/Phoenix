package tiny

// Min yields the smallest value of all provided operands as a string Formula.
//
// see.BaselessCalculation
func Min(operands ...any) Formula[string] {
	return String.Min(operands...)
}

// Min yields the smallest value of all provided operands.
//
// see.BaselessCalculation
func (p Primitive[TOut]) Min(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Min yields the smallest value of all provided operands.
//
// see.BaselessCalculation
func (o Operation[TOut]) Min(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Min yields the smallest value of all provided operands.
//
// see.BaselessCalculation
func (f Formula[TOut]) Min(operands ...any) Formula[TOut] {
}
