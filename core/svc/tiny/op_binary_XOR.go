package tiny

// XOR performs a bitwise logical XOR operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) XOR(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// XOR performs a bitwise logical XOR operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) XOR(b any) Formula[TOut] {
}
