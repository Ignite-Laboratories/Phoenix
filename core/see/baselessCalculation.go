package see

// BaselessCalculation
/*
# " 𝐵𝑎𝑠𝑒𝑙𝑒𝑠𝑠 𝐶𝑎𝑙𝑐𝑢𝑙𝑎𝑡𝑖𝑜𝑛 "

	tl;dr - for "how to use" the calculator, please skip to the 'Performing Arithmetic' section

# Abstract

This is a contextual arithmetic operation where the radix and precision are negotiated at runtime,
rather than at compile time through strict types.  The ultimate goal is to eliminate the stigma
that computers are incapable of handling floating point numbers well.  This stereotype has been perpetuated by
a fundamental feature of the IEEE 754 specification - one that has built the entire world around you, but is also
very easy to resolve: their floats cannot accurately represent the number `0.1`

Baseless calculation does NOT improve the performance of IEEE 754 - nor does it even remotely aim to!  Most systems
don't need to calculate millions of floating point values per second, which is where 𝑡𝑖𝑛𝑦 steps in.  𝑡𝑖𝑛𝑦 acts as an
infinitely precise calculator where all calculation derives from -contextual- precision and base.

𝑡𝑖𝑛𝑦 achieves this by performing all calculations as strings internally and yielding the result in your desired type.

# Bases

𝑡𝑖𝑛𝑦 can operate in literally any unsigned integer base except 0 or 1.  If your context describes those bases,
𝑡𝑖𝑛𝑦 will gracefully default to base₁₀.  For base₂ through base₁₆ placeholders are treated as typically expected, but
for base₁₇ and onwards each placeholder is represented by a multi-character hexadecimal value separated by whitespace
and wrapped in parentheses.

# String Representations

All numeric values in 𝑡𝑖𝑛𝑦 are emitted in the following style:

    // An integer
    42

    // A rational
    42.1234

    // A periodic
    1234.5678‾90

    // An irrational
    ~1234.567890612347571236478123745

    // A based integer
    2#101010

    // A based rational
    7#77.654321076

    // A high-based periodic
    222#(7 AB 3 . 01 6 F0 8 ‾ 9 FB)

    // A based irrational
    ~8#1234.567802612347571236478123745

The non-numeric special characters will -always- appear in the following order, though only some may be present in the emitted value:

  0. "~" - irrationality -  the value was observed as irrational
  1. "#" - radix - the number to the left indicates the radix of the value
  2. "." - decimal - the decimal place breaks the whole and fractional parts
  3. "‾" - periodicity - the digits to the right of the overscore are repeated infinitely

If the output number does not include a radix, base₁₀ is universally implied.

For higher bases each placeholder of the value wrapped in parentheses is represented by the minimum number of hexadecimal
characters necessary, and whitespace is placed between each logical member (including the special characters).

# Performing Arithmetic

Most arithmetic derives from a tiny.Operation - which describes a path to the tiny.Context from which the calculation
should be performed.  The context dictates the radix and precision of all subsequent arithmetic.  To gain a tiny.Operation,
you have four options:

  0. By calling a string default context arithmetic method off of the tiny package
  1. By calling a typed default context arithmetic method off of the tiny package
  2. By creating a default context operation as a composite literal
  3. By casting a std.Path leading to a tiny.Context as a tiny.Operation[TOut]

For example, here's how to add 42 and 8 using all three methods:

	0 - tiny.Add(42, 8) // string result, base₁₀
	1 - tiny.Int8.Add(42,8) // int8 result, base₁₀
	2 - tiny.Operation[complex64]{}.Add(42, 8) // complex64 result, base₁₀
	3 - tiny.Operation[float64](std.Path{"Step0","MyCtx"}).Add(42, 8) // float64 result, contextual base

In the final option, the referenced path must lead to a tiny.Context structure which establishes the basis of calculation
(effectively taking us away from a "baseless" calculator).  If it's unable to yield a tiny.Context
from the prescribed path, 𝑡𝑖𝑛𝑦 will panic as it cannot assume you'll accept it falling back to the default context
and should alert you of the issue.

If you'd like to explicitly create a tiny.Operation referencing the default context, please use an empty path:

	// Both ways reference an atlas.Radix, atlas.Precision, byte context
	op := tiny.Operation[byte](std.Path{})
	op  = tiny.Operation[byte]{}

# Supported Types

𝑡𝑖𝑛𝑦 can parse and emit arithmetic in the following formats:

	// base₁₀ - tiny.Numeric types
	uint, uint8, uint16, uint32, uint64, uintptr
	int, int8, int16, int32, int64
	float32, float64
	complex64, complex128

	// addressable-range base
	[]uint, []uint8, []uint16, []uint32, []uint64, []uintptr
	[]int, []int8, []int16, []int32, []int64

	// base-converted
	*big.Int, *big.Float, *big.Rat
	string

During parsing or emission the radix is implied by the datatype.  Primitive Go types are implicitly treated as base₁₀,
and slices or arrays of them treat their addressable range as the radix - meaning a []byte is a base₂₅₆ value, with
each element representing a single placeholder.  Strings and `big` types are base-converted from their stored value
to the contextual radix.

When using signed types, negative placeholders are parsed as holding their absolute value.

# Inf, NaN, and signed zeros

The IEEE 754 specification describes three unique states outside the interval of finite numbers - '+/- Inf' and 'NaN'.
In addition, a signed zero is an acceptable way of yielding a directional infinity, rather than NaN, when dividing by
zero.  I appreciate and intend to respect and honor those established insights from our ancestors, as they are quite
sharp!  The specification also describes a "signaling" vs "quiet" NaN, used to indicate when something has gone awry
in yielding it - Go does not support this distinction, thus (while I appreciate it) 𝑡𝑖𝑛𝑦 doesn't by extension.

𝑡𝑖𝑛𝑦 will print Inf as an explicit identifier: "∞"

NOTE: 𝑡𝑖𝑛𝑦 only emulates these features of the specification!  No arithmetic is ever performed -while- IEEE 754 encoded (unless in Primitive Mode.)

To provide a signed zero to a calculation, your input type must be a string explicitly carrying the sign with it
or a signed zero IEEE 754 encoded float value.

NOTE: Inf and NaN is only supported when working with string or IEEE 754 float output types, otherwise they will cause panics.

# Advanced Context Features

𝑡𝑖𝑛𝑦 also supports "bounding" a number within a minimum and/or maximum value directly on the tiny.Context.  By
default, the minimum and maximum values are `nil` - which denotes that direction as "unbounded."  If either is
set, and the value crosses the defined threshold, it will either overflow/underflow or "clamp."  When
the value is clamped, it means it "pins" to that boundary and stops.  Otherwise, it will roll over to the opposite
boundary and continue counting.

NOTE: The minimum and maximum values are 'inclusive'

The most important thing I want you to note is that the numbers themselves do NOT carry any information about
their boundaries - the -context- does!  This can be very powerful when working within unique numeric requirements,
such as when emulating restricted color spaces.  Rather than carrying the boundary information along, numbers
are simply filtered through a different context during calculation to yield a valid result.

Another way to think about it is that defining a minimum and maximum establishes a "closed interval," which also is a way
to describe a logical numeric "index."  For example, you can restrict calculation to a byte index by setting the
bounds to a closed interval of [0, 255].

To quickly create a bounded base₁₀ context, you may use the tiny.Bounded method

	tiny.Bounded[float64](0, 42) // Creates a base₁₀ float64 context bounded to [0,42]

# Primitive Mode

𝑡𝑖𝑛𝑦 is not a performant calculator by any means, nor does it claim to be.  However, if you'd like to take advantage
of the formula engine -without- the advanced functionality that 𝑡𝑖𝑛𝑦 offers for radix and precision, you can set your
tiny.Context into "primitive mode" by setting `Primitive` to true.  This does a few things:

  0. Only built-in tiny.Numeric Go types are supported - strings are not
  1. All arithmetic happens within the rules of the primitive Go type you provide
  2. Only minimal implicit type inference and coercion can be performed
  3. 𝑡𝑖𝑛𝑦 will panic whenever you have broken these rules
  4. Floats are calculated while IEEE 754 encoded

This basically means that if you wish to add a `float32` and `float64`, you can - but you cannot add a `complex64`
and `float64`.  While we can do a minimal amount of coercion like this, without strings we're limited to the rules
of Go's built-in type arithmetic.

The primitive Go types are any built-in int, uint, float, or complex type defined by the tiny.Numeric interface.

    NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE
    NOTE                                                                            NOTE
    NOTE  Primitive Mode inherently re-enables the inability to represent "0.1"!!!  NOTE
    NOTE                                                                            NOTE
    NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE NOTE

To quickly create a primitive context (implicitly restricted to base₁₀), you may use the tiny.Primitive method

	tiny.Primitive[complex128]() // Creates a primitive complex128 context

# Overriding the Defaults

If you'd like to override the defaults, you can do so using the following atlas keys:

	atlas.PrecisionMinimum - sets the minimum number of digits to "pretty-print" irrational values to
	atlas.Precision - sets the default number of based placeholders to calculate UP to [256 by default]
	atlas.Radix - sets the default base [base₁₀ by default]

# Living Structures

This was a primer on how to use living structures to establish a default and implicit context for your intelligent
designs to follow.  While it may appear dense, it's really not - when you consider what the following expression implies:

	op := tiny.Operation[int](std.Path{"Some", "Locatable", "Value"})

The operation type IS a path!  This means the structure has no "body" but can still recall its form from a location
stored in memory.  At creation, it's imbued with a predestined path describing its place in the grand index of
thought and creation, and at runtime it must parse whatever emits from that source - as it cannot invoke its own
execution.

Just as WE are "living structures" which must reconcile our own existence upon every invocation =)
*/
type BaselessCalculation struct{}
