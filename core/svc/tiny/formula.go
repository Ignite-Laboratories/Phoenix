package tiny

import "phoenix/core/std"

// A Formula is used to lazily chain contextual operations together.
//
// see.BaselessCalculation
type Formula[TOut any] struct {
	ctx std.Path

	op   Operation[TOut]
	p    Primitive[TOut]
	fn   func(Operation[TOut]) TOut
	fnp  func(Primitive[TOut]) TOut
	name string

	// As is used to convert the yielded result from the underlying Context's type into a supported type.
	//
	// NOTE: This is equivalent to calling 'Equals' on the Formula, just with a different output type.
	//
	// see.BaselessCalculation
	As as[TOut]

	// Variables accesses the named variable members of this Formula.
	//
	// see.BaselessCalculation
	Variables variables[TOut]
}

// To chains in a context switch mid-formula.
//
// see.BaselessCalculation
func (f Formula[TOut]) To(path std.Path, code ...any) Formula[TOut] {

}

// Equals yields the result of the Formula into the Context's type.
//
// NOTE: To yield a different type without affecting the Context, please see the As field.
//
// see.BaselessCalculation
func (f Formula[TOut]) Equals() TOut {
}

// Prove yields the result of the Formula, the LaTeX formula which yielded it, and a LaTeX output describing the proof
// of the steps in generating the formula.
//
// NOTE: This eventually will evolve to support intelligent logical inference and reasoning =)
//
// see.BaselessCalculation
func (f Formula[TOut]) Prove() (result TOut, formula string, proof string) {

}

// String outputs the Formula in LaTeX form.
//
// NOTE: To print the variables that constitute the Formula, please see the Variables field.
//
// see.BaselessCalculation
func (f Formula[TOut]) String() TOut {

}
