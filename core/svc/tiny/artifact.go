package tiny

// An Artifact is an operational side effect.  It differs from a Formula in that it yields a different type (TArt)
// than the Context (TOut), even though its result is generated from an Operation[TOut].  Because of this, the fluent
// chain stops at this point and this describes a 'side effect' of the mathematical formula.
//
// see.BaselessCalculation
type Artifact[TArt any, TOut any] struct {
	op Operation[TOut]
	fn func(Operation[TOut]) TArt
}

func (a Artifact[TArt, TOut]) Operation() Operation[TOut] {
	return a.op
}

func (a Artifact[TArt, TOut]) Equals() TArt {
	return a.fn(a.op)
}

// String outputs the Artifact in LaTeX form.
//
// see.BaselessCalculation
func (a Artifact[TArt, TOut]) String() TOut {

}
