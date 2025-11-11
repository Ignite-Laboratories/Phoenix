package tiny

// AND performs a bitwise logical AND operation between the source and b.
//
// see.BaselessCalculation
func (p Primitive[TOut]) AND(b any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// AND performs a bitwise logical AND operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) AND(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// AND performs a bitwise logical AND operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) AND(b any) Formula[TOut] {
}
