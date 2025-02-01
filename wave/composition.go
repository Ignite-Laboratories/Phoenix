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
			for length := MaxForm; length > 0; length-- {
				wave := c.IdentifyWave(Wave{
					Form:   make([]tiny.Bit, length),
					Period: tiny.Nibble(period),
					Phase:  tiny.Crumb(phase),
				})

				if wave.Frequency >= 4 && len(wave.Form) > 1 {
					fmt.Printf("%v\n", wave)
				}
			}
		}
	}
}

func (c *Composition) ExtractWave(w Wave) {

}

// IdentifyWave takes in partial wave information and projects it onto the timeline while finding its frequency and form.
func (c *Composition) IdentifyWave(w Wave) Wave {
	index := int(w.Phase)

formLoop:
	for index < len(c.Timeline) {
		for formIndex := 0; formIndex < len(w.Form) && index < len(c.Timeline); formIndex++ {
			if w.Frequency == 0 {
				w.Form[formIndex] = c.Timeline[index]
			} else {
				// Check if the timeline matches this occurrence's index
				if c.Timeline[index] != w.Form[formIndex] {
					break formLoop // ...break if not
				}
			}

			// If we fully finish walking a pattern, increment the frequency
			if formIndex == len(w.Form)-1 {
				w.Frequency++
			}

			// Walk forward the wave's periodicity
			index += int(w.Period) + 1
		}
	}

	return w
}
