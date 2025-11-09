package tiny

// XNOR performs a bitwise logical XNOR operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) XNOR(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// XNOR performs a bitwise logical XNOR operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) XNOR(b any) Formula[TOut] {
}
