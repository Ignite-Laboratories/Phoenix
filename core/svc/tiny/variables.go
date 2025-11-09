package tiny

// variables provide a way to retrieve the stored variable members of a Formula.
type variables[TOut any] struct {
	fn func() Formula[TOut]
}

// Reveal returns the named variable members of a Formula.
func (v variables[TOut]) Reveal() []Formula[TOut] {

}

// String prints the named variable members of a Formula in LaTeX.
func (v variables[TOut]) String() string {

}
