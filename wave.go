package main

import "fmt"

// A Wave defines how to project a value across time spatially.
// The Form is a slice of bits up to 4 indices long.
// The Frequency is how many times to repeat the form.
// The Period is how many indices to skip between bits.
type Wave struct {
	Form      []bit  // The bit pattern to project
	Frequency morsel // The number of occurrences to project
	Period    morsel // How much space to apply between occurrences
}

func (w Wave) String() string {
	// ω - Pattern form
	// τ - Period
	// ν - Frequency
	return fmt.Sprintf("%dᴪ %dν %dτ", w.Form, w.Frequency, w.Period)
}
