package tiny

// ShiftLeft shifts the current value "left" by the specified number of placeholders by appending 0s to the end as a string Formula.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func ShiftLeft(places uint) Formula[string] {
	return String.ShiftLeft(places)
}

// ShiftLeft shifts the current value "left" by the specified number of placeholders by appending 0s to the end.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func (p Primitive[TOut]) ShiftLeft(places uint) Formula[TOut] {
	return Formula[TOut]{
		p: p,
		fn: func(op Operation[TOut]) TOut {
			var zero TOut
			return zero
		},
	}
}

// ShiftLeft shifts the current value "left" by the specified number of placeholders by appending 0s to the end.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func (o Operation[TOut]) ShiftLeft(places uint) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// ShiftLeft shifts the current value "left" by the specified number of placeholders by appending 0s to the end.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func (f Formula[TOut]) ShiftLeft(places uint) Formula[TOut] {
}
