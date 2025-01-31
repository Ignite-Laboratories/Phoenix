package main

import "fmt"

// A Wave defines how to project a value across time spatially.
type Wave struct {
	Pattern   []bit // The bit pattern to project
	Frequency int   // The number of occurrences to project
	Phase     int   // How far in the first occurrence should start
	Period    int   // How much space to apply between occurrences
}

func (w Wave) String() string {
	// ω - Waveform
	// ᴪ - Pattern/Particle
	// φ - Phase
	// τ - Period
	// ν - Frequency
	return fmt.Sprintf("·ω - ν%d φ%d τ%d ᴪ%d", w.Pattern, w.Frequency, w.Phase, w.Period)
}
