package main

import (
	"math"
)

func main() {
	ipsum := GetIpsum(5)
	bits := BytesToBits(ipsum)

	UnPaint(bits, []Node{})
}

// UnPaint creates a pathway of steps to recursively remove all 1s from a provided slice of bits
func UnPaint(data []bit, pathway []Node) []Node {
	bestNode := Node{}
	bestCount := CountOnes(data)

	// Node 1: Find the best pattern
PatternLoop:
	for scale := crumb(0); scale < 4; scale++ {
		for frequency := crumb(0); frequency < 4; frequency++ {
			for pattern := nibble(15); pattern < 16; pattern-- {
				node := Node{
					Pattern:   pattern,
					Scale:     scale,
					Frequency: frequency,
				}
				
				synthesized := SynthesizePattern(len(data), node)
				result := XOR(data, synthesized)
				count := CountOnes(result)
				if count < bestCount {
					bestNode = node
					bestCount = count

					// Break out of the entire loop structure when the count is empty
					if count == 0 {
						break PatternLoop
					}
				}
			}
		}
	}
	// Add the best found step to the pathway
	pathway = append(pathway, bestNode)

	// Node 2: Check if the count is empty or if subdivision is impossible
	count := CountOnes(data)
	if count == 0 || len(data) == 1 {
		// ...if so, add a `terminus` value to the pathway and return
		pathway = append(pathway, Node{Scale: 0})
		return pathway
	}

	// Node 3: Subdivide and recurse
	//halves := Subdivide(data)
	//pathway = append(pathway, UnPaint(halves[0], pathway)...)
	//pathway = append(pathway, UnPaint(halves[1], pathway)...)
	return pathway
}

// SynthesizePattern uses the provided Node information to project 1s and 0s across a new slice
// of the provided width
func SynthesizePattern(width int, node Node) []bit {
	//	timeline := make([]bit, width)
	//	startIndex := 0
	//	occurrences := node.Frequency
	//	scaledPattern := Render(node.Pattern, node.Scale)

	//	if node.Frequency == 0 {
	// This is a tiled pattern
	//		occurrences = width
	//	} else {
	// This is a distributed pattern
	//	}

	Render(node.Pattern, float64(width), float64(node.Scale))
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
