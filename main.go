package main

import (
	"math"
)

func main() {
	ipsum := GetIpsum(5)
	bits := BytesToBits(ipsum)

	UnPaint(bits, []Step{})
}

// UnPaint creates a pathway of steps to recursively remove all 1s from a provided slice of bits
func UnPaint(data []bit, pathway []Step) []Step {
	bestStep := Step{}
	bestEnergy := CountOnes(data)

	// Step 1: Find the best pattern
PatternLoop:
	for pattern := nibble(15); pattern < 16; pattern-- {
		for scale := nibble(15); scale <= 1; scale-- { // NOTE: scale '0' is a terminus and ignored
			for frequency := crumb(0); frequency <= 3; frequency++ {
				for offset := morsel(0); offset <= 63; offset++ {
					step := Step{
						Pattern:   pattern,
						Scale:     scale,
						Frequency: frequency,
						Offset:    offset,
					}
					synthesized := SynthesizePattern(len(data), step)
					result := XOR(data, synthesized)
					energy := CountOnes(result)
					if energy < bestEnergy {
						bestStep = step
						bestEnergy = energy

						// Break out of the entire loop structure when energy is first reduced
						if energy == 0 {
							break PatternLoop
						}
					}
				}
			}
		}
	}
	// Add the best found step to the pathway
	pathway = append(pathway, bestStep)

	// Step 2: Check if all the energy is gone or subdivision is impossible
	energy := CountOnes(data)
	if energy == 0 || len(data) == 1 {
		// ...if so, add a `terminus` value to the pathway and return
		pathway = append(pathway, Step{Scale: 0})
		return pathway
	}

	// Step 3: Subdivide and recurse
	//halves := Subdivide(data)
	//pathway = append(pathway, UnPaint(halves[0], pathway)...)
	//pathway = append(pathway, UnPaint(halves[1], pathway)...)
	return pathway
}

// SynthesizePattern uses the provided Step information to project 1s and 0s across a new slice
// of the provided width
func SynthesizePattern(width int, step Step) []bit {
	//	timeline := make([]bit, width)
	//	startIndex := 0
	//	occurrences := step.Frequency
	//	scaledPattern := Render(step.Pattern, step.Scale)

	//	if step.Frequency == 0 {
	// This is a tiled pattern
	//		occurrences = width
	//	} else {
	// This is a distributed pattern
	//	}

	Render(step.Pattern, float64(width), float64(step.Scale))
	return nil
}

// Render takes the provided pattern and scales it to size
func Render(pattern nibble, maxWidth float64, scale float64) []bit {
	totalWidth := int(math.Ceil(4 + ((scale-1)/14)*(maxWidth-4)))
	bitWidth := int(math.Ceil(float64(totalWidth) / 4))
	p := ToBits(pattern)
	result := make([]bit, int(totalWidth))

	patternIndex := 0
	subIndex := 0
	for i := 0; i < totalWidth; i++ {
		if subIndex > bitWidth {
			patternIndex++
			subIndex = 0
		}
		subIndex++

		result[i] = p[patternIndex]
	}

	return result
}
