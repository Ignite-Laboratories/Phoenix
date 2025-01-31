package main

import "fmt"

// A Composition represents a timeline of data and the waves to reconstitute it.
type Composition struct {
	Name     string
	Timeline []bit
	Waves    []Wave
}

// NewComposition starts a new timeline and trims it's ending zeros.
func NewComposition(name string, bits []bit) Composition {
	c := Composition{
		Name:     name,
		Timeline: bits,
		Waves:    []Wave{},
	}
	return c
}

func (c *Composition) String() string {
	str := fmt.Sprintf("----------------------------------------\n")
	str += fmt.Sprintf("Composition: %v [%d]\n", c.Name, len(c.Timeline))
	str += fmt.Sprintf("ω: %v\n", c.Timeline)
	for i := 0; i < len(c.Waves); i++ {
		str += fmt.Sprintf("\tΔ %d: %v\n", i, c.Waves[i])
	}
	str += fmt.Sprintf("----------------------------------------\n")
	return str
}
