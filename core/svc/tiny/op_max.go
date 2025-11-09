package tiny

// Max yields the largest value of all provided operands as a string Formula.
//
// see.BaselessCalculation
func Max(operands ...any) Formula[string] {
	return String.Max(operands...)
}

// Max yields the largest value of all provided operands.
//
// see.BaselessCalculation
func (o Operation[TOut]) Max(operands ...any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Max yields the largest value of all provided operands.
//
// see.BaselessCalculation
func (f Formula[TOut]) Max(operands ...any) Formula[TOut] {
}
