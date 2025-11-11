package tiny

// Add adds each operand to the last in the order provided as a string Formula.
//
// see.BaselessCalculation
func Add(operands ...any) Formula[string] {
	return String.Add(operands...)
}

// Add adds each operand to the last in the order provided.
//
// see.BaselessCalculation
func (p Primitive[TOut]) Add(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Add adds each operand to the last in the order provided.
//
// see.BaselessCalculation
func (o Operation[TOut]) Add(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Add adds each operand to the last in the order provided.
//
// see.BaselessCalculation
func (f Formula[TOut]) Add(operands ...any) Formula[TOut] {
}
