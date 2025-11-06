package tiny

// Context sets the basis for 𝑡𝑖𝑛𝑦 calculations, such as the Radix and Precision.
//
// see.BaselessCalculation
type Context struct {
	// Radix defines the mathematical base of calculation and is only limited by the type - see.Calculation.
	//
	// NOTE: If zero or one, this implicitly means base₁₀.
	//
	// For base₁₇ and on, every placeholder should be represented by minimal hexadecimal characters, separated by whitespace, and wrapped in parentheses.
	//
	// see.BaselessCalculation
	Radix uint

	// Precision defines the number of fractional placeholders to limit calculation to.
	//
	// see.BaselessCalculation
	Precision uint

	// Primitive indicates if all calculation should be performed as primitive Go types and not use the advanced 𝑡𝑖𝑛𝑦 engine.
	//
	// This allows you to use the formula system for lazy chains of calculation with no loss of performance.
	//
	// see.BaselessCalculation
	Primitive bool

	// Clamp indicates if the value can overflow or underflow the defined boundaries, or if it should stop at each extremity.
	//
	// NOTE: This is only factored in when Minimum and/or Maximum are defined and defaults to "false"
	//
	// see.BaselessCalculation
	Clamp bool

	// Minimum bounds the value to exist above the defined value.  The default of `nil` will leave the boundary unbound.
	//
	// see.BaselessCalculation
	Minimum any

	// Maximum bounds the value to exist below the defined value.  The default of `nil` will leave the boundary unbound.
	//
	// see.BaselessCalculation
	Maximum any
}
