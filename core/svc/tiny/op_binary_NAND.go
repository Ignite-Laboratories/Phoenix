package tiny

// NAND performs a bitwise logical NAND operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) NAND(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// NAND performs a bitwise logical NAND operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) NAND(b any) Formula[TOut] {
}
