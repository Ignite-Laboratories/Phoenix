package tiny

// NOT performs a bitwise logical NOT operation on the source operand.
//
// see.BaselessCalculation
func (o Operation[TOut]) NOT() Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// NOT performs a bitwise logical NOT operation on the source operand.
//
// see.BaselessCalculation
func (f Formula[TOut]) NOT() Formula[TOut] {
}
