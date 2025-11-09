package tiny

import "git.ignitelabs.net/janos/core/enum/direction/ordinal"

// Compare calculates the ordinal.Direction of a relative to b as a string Formula.
//
// ordinal.Negative (-1) - a is relatively before b
//
// ordinal.Static (0) - a is relatively the same as b
//
// ordinal.Positive (1) - a is relatively after b
//
// NOTE: Compare does not participate in formulaic chains, as it requires a restriction to a signed type.
//
// see.BaselessCalculation
func Compare(a, b any) Artifact[ordinal.Direction, string] {
	return String.Compare(a, b)
}

// Compare calculates the ordinal.Direction of the source operation's result relative to b.
//
// ordinal.Negative (-1) - a is relatively before b
//
// ordinal.Static (0) - a is relatively the same as b
//
// ordinal.Positive (1) - a is relatively after b
//
// NOTE: Compare yields an Artifact, rather than a Formula, as its output is restricted to an ordinal.Direction (and
// implicitly, signed types.)
//
// see.BaselessCalculation
func (o Operation[TOut]) Compare(a, b any) Artifact[ordinal.Direction, TOut] {
	return Artifact[ordinal.Direction, TOut]{
		op: o,
		fn: func(op Operation[TOut]) ordinal.Direction {

		},
	}
}

// Compare calculates the ordinal.Direction of the source operation's result relative to b.
//
// ordinal.Negative (-1) - a is relatively before b
//
// ordinal.Static (0) - a is relatively the same as b
//
// ordinal.Positive (1) - a is relatively after b
//
// NOTE: Compare yields an Artifact, rather than a Formula, as its output is restricted to an ordinal.Direction (and
// implicitly, signed types.)
//
// see.BaselessCalculation
func (f Formula[TOut]) Compare(b any) Artifact[ordinal.Direction, TOut] {
}
