package main

// A bit represents a single binary value [0 - 1]
type bit byte

const (
	Zero bit = 0
	One  bit = 1
)

// ToBits takes an integer and returns its constituent bits
func ToBits(num int) []bit {
	if num == 0 {
		return []bit{bit(0)}
	}

	var bs []bit
	for num > 0 {
		b := bit(num % 2)  // Get the least significant bit
		bs = append(bs, b) // Append the bit
		num /= 2           // Shift right by dividing by 2
	}

	for left, right := 0, len(bs)-1; left < right; left, right = left+1, right-1 {
		bs[left], bs[right] = bs[right], bs[left]
	}

	return bs
}

// BytesToBits takes a slice of bytes and returns a slice of all of its individual bits
func BytesToBits(data []byte) []bit {
	dataBits := make([]bit, 0, len(data)*8)
	for _, b := range data {
		dataBits = append(dataBits, ToBits(int(b))...)
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
