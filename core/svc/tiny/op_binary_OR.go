package tiny

// OR performs a bitwise logical OR operation between the source and b.
//
// see.BaselessCalculation
func (p Primitive[TOut]) OR(b any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// OR performs a bitwise logical OR operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) OR(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// OR performs a bitwise logical OR operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) OR(b any) Formula[TOut] {
}
