package tiny

// Modulo returns the result of a % b as a string Formula.
//
// see.BaselessCalculation
func Modulo(a, b any) Formula[string] {
	return String.Modulo(a, b)
}

// Modulo returns the result of a % b
//
// see.BaselessCalculation
func (o Operation[TOut]) Modulo(a, b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Modulo returns the result of a % b
//
// see.BaselessCalculation
func (f Formula[TOut]) Modulo(a, b any) Formula[TOut] {
}
