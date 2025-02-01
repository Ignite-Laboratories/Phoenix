package wave

import (
	"Phoenix/tiny"
	"fmt"
)

const MaxForm = 16
const MaxPhase = 256  // tiny.CrumbMax
const MaxPeriod = 256 // tiny.NibbleMax
const MinFrequency = 4
const MaxFrequency = tiny.MorselMax

// A Wave defines how to project a value across time spatially.
// The Form is the repeating pattern, up to 4 indices long.
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
	return fmt.Sprintf("%3dχ\t%3dτ\t%3dφ\t%vƒ", w.Frequency, w.Period, w.Phase, w.Form)
}
