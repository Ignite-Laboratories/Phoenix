package wave

import (
	"Phoenix/tiny"
	"fmt"
)

// A Wave defines how to project a value across time spatially.
// The Form is a slice of bits up to 4 indices long.
// The Period is how many indices to skip between bits. [0-15]
// The Phase is how many indices in to start the projection. [0-3]
// The Frequency is how many times to repeat the form. [0-63]
type Wave struct {
	Form      []tiny.Bit  // The bit pattern to project
	Period    tiny.Nibble // How much space to apply between occurrences
	Phase     tiny.Crumb  // How far in the projection starts
	Frequency tiny.Morsel // The number of occurrences to project
}

func (w Wave) String() string {
	return fmt.Sprintf("%vƒ\t%3dχ\t%3dτ\t%3dφ", w.Form, w.Frequency, w.Period, w.Phase)
}
