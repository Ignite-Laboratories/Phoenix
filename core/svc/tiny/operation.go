package tiny

import "git.ignitelabs.net/janos/core/enum/direction/ordinal"

// A Operation is the source of a calculation chain, beginning with a path to the tiny.Context.
//
// see.BaselessCalculation
type Operation[T any] std.Path

// Compare calculates the ordinal.Direction of a relative to b.
//
// ordinal.Negative (-1) - a is relatively before b
//
// ordinal.Static (0) - a is relatively the same as b
//
// ordinal.Positive (1) - a is relatively after b
//
// see.BaselessCalculation
func (op Operation[TOut]) Compare(a, b any) ordinal.Direction {

}

// Operand simply stores the provided operand and yields a formula.
//
// see.BaselessCalculation
func (op Operation[TOut]) Operand(operand any) Formula[TOut] {

}

// Add adds each operand to the last in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Add(operands ...any) Formula[TOut] {
}

// Subtract subtracts each operand from the last in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Subtract(operands ...any) Formula[TOut] {

}

// Multiply multiplies each operand together in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Multiply(operands ...any) Formula[TOut] {

}

// Divide returns the result of a / b
//
// see.BaselessCalculation
func (op Operation[TOut]) Divide(a, b any) Formula[TOut] {

}

// Modulo returns the result of a % b
//
// see.BaselessCalculation
func (op Operation[TOut]) Modulo(a, b any) Formula[TOut] {

}

// Power returns the result of `base` to the power of `exponent`
//
// see.BaselessCalculation
func (op Operation[TOut]) Power(base, exponent any) Formula[TOut] {

}

// Root returns the result of a root calculation, with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3, 9) = 3 [ [3√9]
//
// see.BaselessCalculation
func (op Operation[TOut]) Root(degree, radicand any) Formula[TOut] {

}

// Factorial returns the result of "operand!"
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func (op Operation[TOut]) Factorial(operand any) Formula[TOut] {

}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func (op Operation[TOut]) Negate(operand any) Formula[TOut] {

}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func (op Operation[TOut]) Absolute(operand any) Formula[TOut] {

}
