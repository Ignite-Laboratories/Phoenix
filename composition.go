package main

// A Composition represents either a timeline of data or the waves to recreate data.
// The RemainingZeros is the final count of zeros to append after reconstitution.
type Composition struct {
	Timeline       []bit
	Waves          []Wave
	RemainingZeros int
}

// NewComposition starts a new timeline and trims it's ending zeros.
func NewComposition(bits []bit) Composition {
	c := Composition{
		Timeline: bits,
		Waves:    []Wave{},
	}
	zeroCount := c.TrimEndZeros()
	c.RemainingZeros = zeroCount
	return c
}

// GetFirst1 finds the index of the first occurrence of 1 in the provided bit slice
func (c *Composition) GetFirst1() int {
	if len(c.Timeline) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(c.Timeline); i++ {
		if (c.Timeline)[i] == One {
			break
		}
		count++
	}
	return count
}

// TrimEndZeros takes a slice of bits and trims it's ending zeros
func (c *Composition) TrimEndZeros() int {
	count := 0
	for i := len(c.Timeline) - 1; i >= 0; i-- {
		if (c.Timeline)[i] == Zero {
			count++
		} else {
			break
		}
	}
	c.Timeline = (c.Timeline)[:len(c.Timeline)-count]
	return count
}

// CountOnes counts the number of 1s in a bit slice
func CountOnes(bits []bit) int {
	count := 0
	for i := 0; i < len(bits); i++ {
		count += int(bits[i])
	}
	return count
}
