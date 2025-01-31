package main

import (
	"Phoenix/tiny"
	"Phoenix/wave"
	"fmt"
)

var maxPatternWidth = 2

func main() {
	bits := tiny.BytesToBits(GetIpsum(1))
	fmt.Println("Source:")
	fmt.Println(bits)
	fmt.Println()

	c := wave.NewComposition("Lorem Ipsum", bits)

	c.UnPaint()
	fmt.Println("Un-Painted:")
	fmt.Println(c.String())

	c.Paint()
	fmt.Println("Re-Painted:")
	fmt.Println(c.String())
}
