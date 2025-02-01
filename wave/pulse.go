package wave

import (
	"Phoenix/tiny"
	"fmt"
)

const MaxBeats = 8
const MaxPhase = tiny.MorselMax
const MaxPeriod = tiny.MorselMax
const MaxFrequency = tiny.NibbleMax

// A Pulse defines how to project beats across time spatially.
// Beats are an ordered slice of variable width pulses.
// The Period is how many indices to skip between beats. [0-63]
// The Phase is how many indices in to start the beats. [0-63]
// The Frequency is how many times to repeat the beats. [0-15]
type Pulse struct {
	Beats     []Beat
	Period    tiny.Morsel // How much space to apply between occurrences
	Phase     tiny.Morsel // How far in the projection starts
	Frequency tiny.Nibble // The number of occurrences to project
}

func (w Pulse) String() string {
	return fmt.Sprintf("%3dχ\t%3dτ\t%3dφ\t%vƒ", w.Frequency, w.Period, w.Phase, w.Beats)
}
