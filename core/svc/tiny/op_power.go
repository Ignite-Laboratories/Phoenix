package tiny

// Power returns the result of `base` to the power of `exponent` as a string Formula.
//
// see.BaselessCalculation
func Power(base, exponent any) Formula[string] {
	return String.Power(base, exponent)
}

// Power returns the result of `base` to the power of `exponent`
//
// see.BaselessCalculation
func (o Operation[TOut]) Power(base, exponent any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// Power returns the result of `base` to the power of `exponent`
//
// see.BaselessCalculation
func (f Formula[TOut]) Power(base, exponent any) Formula[TOut] {
}
