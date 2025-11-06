package tiny

import "git.ignitelabs.net/janos/core/enum/direction/ordinal"

// A Operation is the source of a calculation chain, beginning with a path to the tiny.Context.
//
// see.BaselessCalculation
type Operation[T any] std.Path

// Add adds each operand to the last in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Add(operands ...any) TOut {
}

// Subtract subtracts each operand from the last in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Subtract(operands ...any) TOut {

}

// Multiply multiplies each operand together in the order provided.
//
// see.BaselessCalculation
func (op Operation[TOut]) Multiply(operands ...any) TOut {

}

// Divide returns the result of a / b
//
// see.BaselessCalculation
func (op Operation[TOut]) Divide(a, b any) TOut {

}

// Modulo returns the result of a % b
//
// see.BaselessCalculation
func (op Operation[TOut]) Modulo(a, b any) TOut {

}

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

// Power returns the result of `base` to the power of `exponent`
//
// see.BaselessCalculation
func (op Operation[TOut]) Power(base, exponent any) TOut {

}

// Root returns the result of a root calculation, with the `degree` (or index) being outside the radical sign and the `radicand` being inside.
//
// For Example:
//
//	Root(2, 144) = 12 [2√12]
//	Root(3, 9) = 3 [ [3√9]
//
// see.BaselessCalculation
func (op Operation[TOut]) Root(degree, radicand any) TOut {

}

// Factorial returns the result of "operand!"
//
// For Example:
//
//	Factorial(5) = 120 [ 5⋅4⋅3⋅2⋅1 ]
//
// see.BaselessCalculation
func (op Operation[TOut]) Factorial(operand any) TOut {

}

// Negate returns the result of "-operand"
//
// see.BaselessCalculation
func (op Operation[TOut]) Negate(operand any) TOut {

}

// Absolute returns the result of "|operand|"
//
// see.BaselessCalculation
func (op Operation[TOut]) Absolute(operand any) TOut {

}
