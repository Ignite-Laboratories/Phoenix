package wave

import (
	"Phoenix/tiny"
	"fmt"
)

// A Beat is [0] or [1] projected [1-8] times
type Beat struct {
	Value    tiny.Bit
	Duration tiny.Note
}

func (b Beat) String() string {
	bits := b.Expand()
	str := fmt.Sprintf("%v\n", bits)
	return str
}

// Expand takes a beat and expands it to the described size
func (b Beat) Expand() []tiny.Bit {
	bits := make([]tiny.Bit, 0)
	for i := 0; i < int(b.Duration); i++ {
		bits = append(bits, b.Value)
	}
	return bits
}
