package main

type BITS interface {
	bit | crumb | nibble | morsel | byte
}

// A bit represents a single binary value [0 - 1]
type bit byte

// A crumb represents two bits [0 - 3]
type crumb byte

// A nibble represents four bits [0 - 15]
type nibble byte

// A morsel represents six bits [0 - 63 or -32 - 31]
type morsel byte

// A TilePattern defines how to synthesize a set of 1s and 0s across an arbitrarily long timeline
//
// A Scale of 0 is a `terminus` value and n is n:1 scale - i.e. for every bit it's emitted 1-15 wide.
// A Period of 0 is no space between tiles.  For each step up, one tile is removed and the remaining are distributed.
// The Phase is the number of indices to push the pattern through the reduced space.
type TilePattern struct {
	Scale   nibble
	Period  int
	Phase   int
	Pattern BitPattern
}

// A BitPattern defines the source of 1s and 0s from which to synthesize a larger structure.
// Value represents the pattern of 16 1s and 0s
// Trim represents how much the width of the pattern was reduced
// Phase represents how much the pattern is phased across the reduced space
type BitPattern struct {
	Value [16]bit
	Trim  nibble
	Phase nibble
}

// CountOnes counts the number of 1s in a bit slice
func CountOnes(data []bit) int {
	count := 0
	for i := 0; i < len(data); i++ {
		count += int(data[i])
	}
	return count
}

// Subdivide splits a slice into two halves using integer math
func Subdivide(data []bit) [][]bit {
	pivot := len(data) / 2 // This always rounds downward
	return [][]bit{data[:pivot], data[pivot:]}
}

// XOR takes two equal length slices, XORs their respective indices, and returns the results
func XOR(a []bit, b []bit) []bit {
	result := make([]bit, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result
}

// ToBits takes in a BITS type and returns a slice of its constituent bits
func ToBits[T BITS](value T) []bit {
	n := int(value)
	var bits []bit
	bitSize := 0
	switch any(value).(type) {
	case crumb:
		bitSize = 2
	case nibble:
		bitSize = 4
	case morsel:
		bitSize = 6
	case byte:
		bitSize = 8
	}

	if n < 0 {
		n = n & ((1 << bitSize) - 1)
	}

	for i := 0; i < bitSize; i++ {
		bits = append([]bit{bit(n & 1)}, bits...)
		n >>= 1
	}

	return bits
}

// FromBits takes in a BITS type and returns a slice of its constituent bits
func FromBits[T BITS](bits []bit) T {
	bitSize := byte(0)
	switch any(new(T)).(type) {
	case crumb:
		bitSize = 2
	case nibble:
		bitSize = 4
	case morsel:
		bitSize = 6
	case byte:
		bitSize = 8
	}

	result := byte(0)
	for i, b := range bits {
		result |= byte(b) << ((bitSize - byte(1)) - byte(i))
	}

	return T(result)
}

// BytesToBits takes a slice of bytes and returns a slice of all of its individual bits
func BytesToBits(data []byte) []bit {
	dataBits := make([]bit, 0, len(data)*8)
	for _, b := range data {
		dataBits = append(dataBits, ToBits(b)...)
	}
	return dataBits
}

// GetIpsum returns the provided number of paragraphs of 'Lorem ipsum' for sample data
func GetIpsum(paragraphs int) []byte {
	ipsum := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque imperdiet libero eu neque facilisis, ac pretium nisi dignissim. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.`
	toReturn := ipsum
	for i := 0; i < paragraphs; i++ {
		toReturn += ipsum
	}
	return []byte(toReturn)
}
