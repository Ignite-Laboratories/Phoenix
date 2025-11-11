package tiny

import (
	"phoenix/core/std"

	"git.ignitelabs.net/janos/core/sys/atlas"
)

/**

Primitive Mode

*/

// Numeric is a type-constraint defining the standard Go numeric types for use in a Primitive Context.
type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128
}

// Primitive is the default primitive Context type.  If you'd like to perform primitive calculation, please
// start your chain from this type.
type Primitive[TOut Numeric] struct{}

/**

Precision

*/

func Precision[TOut any](precision uint) Operation[TOut] {

	// TODO: Store a new precision context in std.Memory

	panic("not implemented yet.")
}

/**

Base

*/

type base struct {
	radix     uint
	precision uint
}

var baseCache map[base]std.Path

func Base[TOut any](radix uint, precision ...uint) Operation[TOut] {
	p := atlas.Precision
	if len(precision) > 0 {
		p = precision[0]
	}

	b := base{
		radix:     radix,
		precision: p,
	}
	if path, exists := baseCache[b]; exists {
		return Operation[TOut](path)
	}

	// TODO: Store a new context in std.Memory

	panic("not implemented yet.")
}

/**

Bounded

*/

func Bounded[TOut any](minimum, maximum TOut, precision ...uint) Operation[TOut] {
	p := atlas.Precision
	if len(precision) > 0 {
		p = precision[0]
	}

	// TODO: Store a new bounded context in std.Memory

	panic("not implemented yet.")
}
