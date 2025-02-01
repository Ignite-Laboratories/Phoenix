package tiny

// A Bit represents a single binary value [0 - 1]
type Bit byte

const Zero Bit = 0
const One Bit = 1

// A Crumb represents two binary values [0-3]
type Crumb byte

const CrumbMax = 4

// A Note represents three binary values [0-7]
type Note byte

const NoteMax = 8

// A Nibble represents four binary values [0-15]
type Nibble byte

const NibbleMax = 16

// A Morsel represents six binary values [0-63]
type Morsel byte

const MorselMax = 64
const ByteMax = 256

// ToBits takes an integer and returns its constituent bits
func ToBits(value int) []Bit {
	if value == 0 {
		return []Bit{Bit(0)}
	}

	var bits []Bit
	for value > 0 {
		bit := Bit(value % 2)    // Get the least significant Bit
		bits = append(bits, bit) // Append the Bit
		value /= 2               // Shift right by dividing by 2
	}

	// Reverse the slice
	for left, right := 0, len(bits)-1; left < right; left, right = left+1, right-1 {
		bits[left], bits[right] = bits[right], bits[left]
	}

	return bits
}

// BytesToBits takes a slice of bytes and returns a slice of all of its individual bits
func BytesToBits(data []byte) []Bit {
	bits := make([]Bit, 0, len(data)*8)
	for _, b := range data {
		bits = append(bits, ToBits(int(b))...)
	}
	return bits
}
