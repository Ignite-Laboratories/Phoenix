package main

import "fmt"

func main() {
	composition := NewComposition(BytesToBits(GetIpsum(1)))
	composition.UnPaint()
}

func (c *Composition) UnPaint() {
	// First, we find the first 1 and set the wave's 'phase' [φ]
	bestWave := Wave{
		Phase: c.GetFirst1(),
	}

	bestResult := make([]bit, len(c.Timeline))
	bestOnes := CountOnes(c.Timeline)
	c.TrimEndZeros()

	// Second, we find the largest size that could potentially repeat
	maxPatternWidth := len(c.Timeline) / 16

	for patternWidth := maxPatternWidth; patternWidth > 0; patternWidth-- { // Walk the biggest pattern down to the smallest [1 bit wide]
		maxFrequency := len(c.Timeline) / patternWidth

		bestFrequency := CountOnes(c.Timeline)
	frequencyLoop:
		for frequency := maxFrequency; frequency > 1; frequency-- { // Walk the frequencies of that pattern down [2 minimum]
			negativeSpace := len(c.Timeline) - (patternWidth * frequency) // Calculate the amount of negative space from this pattern/frequency combination

			for period := 0; period < negativeSpace; period++ { // Walk the negative space upward as periodicity
				w := bestWave
				w.Frequency = frequency
				w.Pattern = c.Timeline[:patternWidth]
				w.Period = period

				result := c.SynthesizeWave(w)
				ones := CountOnes(result)

				if ones < bestOnes {
					bestOnes = ones
					bestWave = w
					bestResult = result
				}
			}

			// If the amount of ones has already reduced at a high frequency, walking the remaining
			// frequencies would ALWAYS increase the count -> pointless calculations.
			if CountOnes(bestResult) < bestFrequency {
				break frequencyLoop
			}
		}
	}

	c.Timeline = bestResult[bestWave.Phase+len(bestWave.Pattern):]
	c.Waves = append(c.Waves, bestWave)

	// Keep recursing until the final one, as it would never reoccur
	//	if len(c.Timeline) > 1 {
	//		c.UnPaint()
	//	}
}

// SynthesizeWave takes a composition and projects the provided wave across it, returning the XORd result.
// Use this to test how a wave de-energizes a composition for comparison.
func (c *Composition) SynthesizeWave(wave Wave) []bit {
	fmt.Println(wave)
	return nil
}
