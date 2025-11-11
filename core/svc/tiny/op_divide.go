package tiny

// Divide returns the result of a / b as a string Formula.
//
// see.BaselessCalculation
func Divide(a, b any) Formula[string] {
	return String.Divide(a, b)
}

// Divide returns the result of a / b
//
// see.BaselessCalculation
func (p Primitive[TOut]) Divide(a, b any) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// Divide returns the result of a / b
//
// see.BaselessCalculation
func (o Operation[TOut]) Divide(a, b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Divide returns the result of a / b
//
// see.BaselessCalculation
func (f Formula[TOut]) Divide(a, b any) Formula[TOut] {
}
