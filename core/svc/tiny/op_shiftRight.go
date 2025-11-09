package tiny

// ShiftRight shifts the current value "right" by the specified number of placeholders by eliminating zeros from
// the end or moving the decimal place leftwards as a string Formula
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func ShiftRight(places uint) Formula[string] {
	return String.ShiftLeft(places)
}

// ShiftRight shifts the current value "right" by the specified number of placeholders by eliminating zeros from
// the end or moving the decimal place leftwards.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func (o Operation[TOut]) ShiftRight(placeholders uint) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// ShiftRight shifts the current value "right" by the specified number of placeholders by eliminating zeros from
// the end or moving the decimal place leftwards.
//
// NOTE: This is base-agnostic - to perform this as a standard bitwise shift, please recall a base₂ Context.
//
// see.BaselessCalculation
func (f Formula[TOut]) ShiftRight(placeholders uint) Formula[TOut] {
}