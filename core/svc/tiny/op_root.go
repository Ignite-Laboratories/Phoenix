package tiny

// Root returns the result of a root calculation as a string Formula, with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3, 9) = 3 [ [3√9]
//
// see.BaselessCalculation
func Root(degree, radicand any) Formula[string] {
	return String.Root(degree, radicand)
}

// Root returns the result of a root calculation, with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3,   9) =  3 [3√9]
//
// see.BaselessCalculation
func (o Operation[TOut]) Root(degree, radicand any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Root returns the result of a root calculation, with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3,   9) =  3 [3√9]
//
// see.BaselessCalculation
func (f Formula[TOut]) Root(degree, radicand any) Formula[TOut] {
}
