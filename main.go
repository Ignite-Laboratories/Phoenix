package main

import "fmt"

func main() {
	composition := NewComposition(BytesToBits(GetIpsum(1)))
	composition.UnPaint()
}

func (c *Composition) UnPaint() {
	// First, we find the first 1 and store it as the wave's 'phase' [φ]
	// We don't walk inward beyond the first relevant 1 as all are important in the process.
	bestWave := Wave{
		Phase: c.GetFirst1(),
	}

	// Setup our 'best' variables
	bestResult := make([]bit, len(c.Timeline))
	bestOnes := CountOnes(c.Timeline)

	// Trim the end of the timeline as it's irrelevant at this point in the process
	c.TrimEndZeros()

	// Find the largest size that could potentially repeat (half the width of the data)
	maxPatternWidth := len(c.Timeline) / 2

	// Walk the biggest pattern down to the smallest as 'pattern' [ᴪ] [1 bit wide]
	for patternWidth := maxPatternWidth; patternWidth > 0; patternWidth-- {

		// Walk the frequencies of that pattern down as 'frequency' [ν] [2 minimum]
		maxFrequency := len(c.Timeline) / patternWidth
		bestFrequency := CountOnes(c.Timeline)
	FrequencyLoop:
		for frequency := maxFrequency; frequency > 1; frequency-- {

			// Calculate the amount of negative space from this pattern/frequency combination
			negativeSpace := len(c.Timeline) - (patternWidth * frequency)

			// Walk the negative space upward as 'periodicity' [τ]
			for period := 0; period < negativeSpace; period++ {
				w := bestWave
				w.Frequency = frequency
				w.Pattern = c.Timeline[:patternWidth]
				w.Period = period

				testResult := c.SynthesizeWave(w)
				ones := CountOnes(testResult)

				if ones < bestOnes {
					bestOnes = ones
					bestWave = w
					bestResult = testResult
				}
			}

			// If the amount of ones has already reduced at a high frequency, walking the remaining
			// frequencies would ALWAYS increase the count -> pointless calculations.
			if CountOnes(bestResult) < bestFrequency {
				break FrequencyLoop
			}
		}
	}

	// Trim the timeline by the phase + best found pattern width
	c.Timeline = bestResult[bestWave.Phase+len(bestWave.Pattern):]
	// Store the best found wave information
	c.Waves = append(c.Waves, bestWave)

	// Recurse if there is still at least 2 ones left on the timeline
	ones := CountOnes(c.Timeline)
	if ones > 1 {
		c.UnPaint()
	}
}

// SynthesizeWave takes a composition and projects the provided wave across it, returning the XORd result.
// Use this to test how a wave de-energizes a composition for comparison.
func (c *Composition) SynthesizeWave(wave Wave) []bit {
	fmt.Println(wave)
	return nil
}
