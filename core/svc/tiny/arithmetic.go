package tiny

import "git.ignitelabs.net/janos/core/enum/direction/ordinal"

// Add adds each operand to the last in the order provided as an Operation[string].
//
// see.BaselessCalculation
func Add(operands ...any) string {
	return String.Add(operands...)
}

// Subtract subtracts each operand from the last in the order provided as an Operation[string].
//
// see.BaselessCalculation
func Subtract(operands ...any) string {
	return String.Subtract(operands...)
}

// Multiply multiplies each operand together in the order provided as an Operation[string].
//
// see.BaselessCalculation
func Multiply(operands ...any) string {
	return String.Multiply(operands...)
}

// Divide returns the result of a / b as an Operation[string].
//
// see.BaselessCalculation
func Divide(a, b any) string {
	return String.Divide(a, b)
}

// Modulo returns the result of a % b as an Operation[string].
//
// see.BaselessCalculation
func Modulo(a, b any) string {
	return String.Modulo(a, b)
}

// Compare calculates the ordinal.Direction of a relative to b as an Operation[string].
//
// ordinal.Negative (-1) - a is relatively before b
//
// ordinal.Static (0) - a is relatively the same as b
//
// ordinal.Positive (1) - a is relatively after b
//
// see.BaselessCalculation
func Compare(a, b any) ordinal.Direction {
	return String.Compare(a, b)
}

// Power returns the result of `base` to the power of `exponent` as an Operation[string].
//
// see.BaselessCalculation
func Power(base, exponent any) string {
	return String.Power(base, exponent)
}

// Root returns the result of a root calculation as an Operation[string], with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3, 9) = 3 [ [3√9]
//
// see.BaselessCalculation
func Root(degree, radicand any) string {
	return String.Root(degree, radicand)
}

// Factorial returns the result of "operand!" as an Operation[string].
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func Factorial(operand any) string {
	return String.Factorial(operand)
}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func Negate(operand any) TOut {
	return String.Negate(operand)
}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func Absolute(operand any) TOut {
	return String.Absolute(operand)
}
