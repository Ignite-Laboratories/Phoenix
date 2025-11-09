package tiny

// NOR performs a bitwise logical NOR operation between the source and b.
//
// see.BaselessCalculation
func (o Operation[TOut]) NOR(b any) Formula[TOut] {
	return Formula[TOut]{
		op: o,
		fn: func(op Operation[TOut]) TOut {

		},
	}
}

// NOR performs a bitwise logical NOR operation between the source and b.
//
// see.BaselessCalculation
func (f Formula[TOut]) NOR(b any) Formula[TOut] {
}
