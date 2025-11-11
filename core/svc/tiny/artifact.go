package tiny

import "phoenix/core/std"

// An Artifact is an operational side effect.  It differs from a Formula in that it yields a different type (TArt)
// than the Context (TOut).  Because of this, the fluent chain stops at this point, and this describes a 'side effect'
// of the mathematical formula.
//
// see.BaselessCalculation
type Artifact[TArt any, TOut any] struct {
	ctx std.Path

	op  Operation[TOut]
	p   Primitive[TOut]
	fn  func(Operation[TOut]) TArt
	fnp func(Primitive[TOut]) TArt
}

// To chains in a context switch mid-formula.
//
// see.BaselessCalculation
func (a Artifact[TArt, TOut]) To(path std.Path, code ...any) Artifact[TArt, TOut] {

}

// Prove yields the result of the Artifact, the LaTeX formula which yielded it, and a LaTeX output describing the proof
// of the steps in generating the formula.
//
// NOTE: This eventually will evolve to support intelligent logical inference and reasoning =)
//
// see.BaselessCalculation
func (a Artifact[TArt, TOut]) Prove() (result TArt, formula string, proof string) {

}

func (a Artifact[TArt, TOut]) Equals() TArt {
	ctx, err := std.Recall[Context](a.ctx)
	if err == nil {
		if ctx.Primitive {
			return a.fnp(a.p)
		}
		return a.fn(a.op)
	}
	panic(err)
}

// String outputs the Artifact in LaTeX form.
//
// see.BaselessCalculation
func (a Artifact[TArt, TOut]) String() TOut {
	panic("LaTeX support is not yet implemented.")
}
