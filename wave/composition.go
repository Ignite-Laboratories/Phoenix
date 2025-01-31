package wave

import (
	"Phoenix/tiny"
	"fmt"
)

// A Composition represents a timeline of data and the waves to reconstitute it.
type Composition struct {
	Name     string
	Timeline []tiny.Bit
	Waves    []Wave
}

// NewComposition starts a new timeline and trims it's ending zeros.
func NewComposition(name string, bits []tiny.Bit) Composition {
	c := Composition{
		Name:     name,
		Timeline: bits,
		Waves:    []Wave{},
	}
	return c
}

func (c *Composition) String() string {
	str := fmt.Sprintf("\"%v\" [%d]\n", c.Name, len(c.Timeline))
	str += fmt.Sprintf("ω: %v\n", c.Timeline)
	for i := 0; i < len(c.Waves); i++ {
		str += fmt.Sprintf("\t%3dΔ - %v\n", i, c.Waves[i])
	}
	return str
}

func (c *Composition) Paint() {
	for i := len(c.Waves) - 1; i >= 0; i-- {
		// TODO: Inject each wave back onto the timeline in reverse
	}
}

func (c *Composition) UnPaint() {
	for phase := 0; phase < MaxPhase; phase++ {
		for period := 0; period < MaxPeriod; period++ {
			for frequency := MinFrequency; frequency < MaxFrequency; frequency++ {

			}
		}
	}
}
