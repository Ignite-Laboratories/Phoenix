package main

import "fmt"

func main() {
	ipsum := GetIpsum(5)
	bits := BytesToBits(ipsum)[:88]
	fmt.Println(bits)
	timeline, pathway := UnPaint(bits, []TilePattern{}, 7)
	fmt.Println(timeline)
	fmt.Println(pathway)
}

// UnPaint creates a pathway of steps to recursively remove all 1s from a provided slice of bits
func UnPaint(data []bit, pathway []TilePattern, threshold int) ([]bit, []TilePattern) {
	bestXOR := data
	bestNode := TilePattern{}
	bestCount := CountOnes(data)

	// Step 1: Find the best pattern
PatternLoop:
	for frequency := crumb(0); frequency < 4; frequency++ {
		for scale := crumb(3); scale > 0; scale-- { // Scale '0' is a terminus and ignored during synthesis
			for pattern := nibble(15); pattern < 16; pattern-- {
				currentNode := TilePattern{
					Pattern:   pattern,
					Scale:     scale,
					Frequency: frequency,
				}

				// Synthesize a pattern to XOR with
				synthesized := SynthesizePattern(len(data), currentNode)

				// Perform the XOR operation to a temporary slice
				temp := XOR(data, synthesized)

				// Check if the delta has improved...
				count := CountOnes(temp)
				if count < bestCount {
					// ...save off the better match
					bestNode = currentNode
					bestCount = count
					bestXOR = temp

					// ...and break out of the entire loop structure when the count is empty
					if count == 0 {
						break PatternLoop
					}
				}
			}
		}
	}
	// Add the best found step to the pathway
	pathway = append(pathway, bestNode)

	// Step 2: Check if the count is empty or if subdivision is impossible
	count := CountOnes(bestXOR)
	if count <= threshold || len(bestXOR) == 1 {
		// ...if so, add a `terminus` value to the pathway and return
		pathway = append(pathway, TilePattern{Scale: 0})
		return data, pathway
	}

	// Step 3: Subdivide and recurse
	halves := Subdivide(bestXOR)
	halfA, pathB := UnPaint(halves[0], pathway, threshold)
	halfB, pathB := UnPaint(halves[1], pathway, threshold)
	pathway = append(pathway, pathB...)
	pathway = append(pathway, pathB...)

	result := make([]bit, 0)
	result = append(result, halfA...)
	result = append(result, halfB...)

	return result, pathway
}

// SynthesizePattern uses the provided TilePattern information to project 1s and 0s across a new slice
// of the provided width.  To do this, we build a slice called the 'timeline' and then tile the
// pattern across it according to the TilePattern values.  The pattern is always placed in the center
// of each tile.
func SynthesizePattern(width int, node TilePattern) []bit {
	timeline := make([]bit, width)
	pattern := ToBits(node.Pattern)
	patternWidth := len(pattern)

	// First, determine how many indexes each bit of the scaled pattern covers
	bitWidth := 0
	switch node.Scale {
	case 1:
		bitWidth = 1
	case 2:
		bitWidth = 2
	case 3:
		bitWidth = 4
	}

	// Second, determine how wide the tile is and the inset for each instance
	scaledWidth := bitWidth * patternWidth
	tileWidth := scaledWidth
	inset := 0
	switch node.Frequency {
	case 1:
		tileWidth = width / 4
		inset = (tileWidth / 2) - (scaledWidth / 2)
	case 2:
		tileWidth = width / 2
		inset = (tileWidth / 2) - (scaledWidth / 2)
	case 3:
		tileWidth = width
		inset = (tileWidth / 2) - (scaledWidth / 2)
	}
	if inset < 0 {
		inset = 0
	}

	// Lastly, walk the tiles and inject a single pattern per instance
TileLoop:
	for i := 0; i < width; i += tileWidth {
		for patternI := 0; patternI < 4; patternI++ {
			for bitI := 0; bitI < bitWidth; bitI++ {
				index := i + inset + (patternI * bitWidth) + bitI
				if index >= len(timeline) {
					break TileLoop
				}
				timeline[index] = pattern[patternI]
			}
		}
	}

	return timeline
}
