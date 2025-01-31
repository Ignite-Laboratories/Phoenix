package tiny

// A Bit represents a single binary value [0 - 1]
type Bit byte

const Zero Bit = 0
const One Bit = 1

// A Crumb represents two binary values [0-3]
type Crumb byte

const CrumbMax Crumb = 4

// A Nibble represents four binary values [0-15]
type Nibble byte

const NibbleMax Nibble = 16

// A Morsel represents six binary values [0-63]
type Morsel byte

const MorselMax Morsel = 64

// ToBits takes an integer and returns its constituent bits
func ToBits(num int) []Bit {
	if num == 0 {
		return []Bit{Bit(0)}
	}

	var bs []Bit
	for num > 0 {
		b := Bit(num % 2)  // Get the least significant Bit
		bs = append(bs, b) // Append the Bit
		num /= 2           // Shift right by dividing by 2
	}

	for left, right := 0, len(bs)-1; left < right; left, right = left+1, right-1 {
		bs[left], bs[right] = bs[right], bs[left]
	}

	return bs
}

// BytesToBits takes a slice of bytes and returns a slice of all of its individual bits
func BytesToBits(data []byte) []Bit {
	dataBits := make([]Bit, 0, len(data)*8)
	for _, b := range data {
		dataBits = append(dataBits, ToBits(int(b))...)
	}
	return dataBits
}
