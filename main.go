package main

import "fmt"

var maxPatternWidth = 2

func main() {
	composition := NewComposition("Lorem Ipsum", BytesToBits(GetIpsum(1)))
	composition.UnPaint()
}

func (c *Composition) UnPaint() {
	c.Waves = append(c.Waves, Wave{
		Form:      []bit{Zero, One, One, Zero},
		Frequency: 42,
		Period:    5,
	})
	c.Waves = append(c.Waves, Wave{
		Form:      []bit{One, One, Zero, Zero},
		Frequency: 31,
		Period:    3,
	})
	c.Waves = append(c.Waves, Wave{
		Form:      []bit{Zero, One, One, One},
		Frequency: 2,
		Period:    15,
	})
	c.Waves = append(c.Waves, Wave{
		Form:      []bit{One, Zero, One, Zero},
		Frequency: 83,
		Period:    7,
	})
	fmt.Println(c)
}
